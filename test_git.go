package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	AppRootDir = "/home/runner/app"
)

type gitCmdParams struct {
	RemoteName   string `json:"remote_name"`
	RemoteBranch string `json:"remote_branch"`
	Message      string `json:"message"`
	File         string `json:"file"`
	LocalBranch  string `json:"local_branch"`
	Username     string `json:"user_name"`
	Token        string `json:"token"`
}

type BaseGitCmd struct {
	CmdType    string `json:"cmd_type"`
	CmdContent string `json:"cmd_content"`
}

const (
	GitCmdTypeAddCommit = "git_add_commit"
)

// git 命令处理handler
type gitCMdHandler func(cmdType string, cmdContent string) (any, error)

// git 处理handler清单
var gitCmdHandlerFactory map[string]gitCMdHandler

func init() {
	baseGitCmd := BaseGitCmd{}
	gitCmdHandlerFactory = map[string]gitCMdHandler{
		GitCmdTypeAddCommit: baseGitCmd.gitCmdAddCommitHandler,
	}
}

// Handler git 命令处理统一入口
func (c *BaseGitCmd) Handler() (any, error) {
	fmt.Printf("gitCmd:git-cmd-handler-begin, cmd_type: %s, cmd_content: %s\n", c.CmdType, c.CmdContent)
	if handler, ok := gitCmdHandlerFactory[c.CmdType]; ok {
		res, err := handler(c.CmdType, c.CmdContent)
		if err != nil {
			fmt.Printf("gitCmd:git-cmd-handler-end, cmd_type: %s, cmd_content: %s\n", c.CmdType, c.CmdContent)
			return "", err
		}

		return res, nil
	}

	return "", errors.New("unknown git cmd")
}

func (c *BaseGitCmd) getCurrentBranch(path string) (string, error) {
	// 保存当前分支
	currentBranch, err := c.execGitCmd(path, "git", "branch", "--show-current")
	if err != nil {
		return "", err
	}
	currentBranch = strings.TrimSuffix(currentBranch, "\n")
	currentBranch = strings.TrimSpace(currentBranch)

	return currentBranch, nil
}

func (c *BaseGitCmd) fetch(path string, cmd *gitCmdParams) (string, error) {
	if cmd.RemoteName == "" {
		return "", errors.New("gitCmdFetch-remote-name is null")
	}

	fmt.Printf("git-cmd-handler-gitCmdFetch, cmd: %+v\n", cmd)

	fetchOut, err1 := c.execGitCmd(path, "git", "fetch", cmd.RemoteName, cmd.RemoteBranch)
	if err1 != nil {
		fmt.Printf("git-cmd-handler-gitCmdFetch-fail, cmd: %+v, beforeOut: %s\n", cmd, fetchOut)
		return "", err1
	}

	fmt.Printf("git-cmd-handler-gitCmdFetch-success, cmd: %+v, currentBranch: %s\n", cmd, fetchOut)

	return fetchOut, nil
}

// reset
func (c *BaseGitCmd) reset(cmd *gitCmdParams, path string) (string, error) {
	fmt.Printf("git-cmd-handler-gitCmdFetchAndReset-reset, cmd: %+v\n", cmd)
	// reset
	out, err := c.execGitCmd(path, "git", "reset", "--hard", fmt.Sprintf("%s/%s", cmd.RemoteName, cmd.RemoteBranch))
	if err != nil {
		return "", err
	}

	fmt.Printf("git-cmd-handler-gitCmdFetchAndReset-reset, cmd: %s, err: %+v, out: %s\n", cmd, err, out)
	return out, nil
}

// 执行git命令
func (c *BaseGitCmd) execGitCmd(dir string, command string, args ...string) (string, error) {
	arg := []string{"bash", "-c", command, strings.Join(args, " ")}
	argStr := strings.Join(arg[2:], " ")
	cmd := exec.Command(arg[0], arg[1], argStr)
	cmd.Dir = dir

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	fmt.Printf("git-cmd-handler-execGitCmd, dir: %s, command: %s, args: %+v, out: %s, err: %v\n", dir, command, args, out.String(), err)

	return out.String(), err
}

// 执行 cmd 命令
func (c *BaseGitCmd) execCmd(dir string, cmd *exec.Cmd) (string, error) {
	cmd.Dir = dir

	// 检查目录是否存在，存在则设置 PATH
	if _, err := os.Stat("/home/runner/app/node_modules/.bin"); os.IsNotExist(err) {
		cmd.Env = os.Environ()
	} else {
		cmd.Env = append(os.Environ(), "PATH="+os.Getenv("PATH")+":/home/runner/app/node_modules/.bin")
	}

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	fmt.Printf("git-cmd-handler-execGitCommitCmd, dir: %s, out: %s\n", dir, out.String())

	return out.String(), err
}

// gitCmdAddCommitHandler: addAndCommit
func (c *BaseGitCmd) gitCmdAddCommitHandler(cmdType string, cmdContent string) (any, error) {
	addDto := &gitDTO{}

	// 计时
	startTime := time.Now()
	defer func() {
		fmt.Printf("gitCmdAddCommitHandler took %v\n", time.Since(startTime))
	}()

	fmt.Printf("git-cmd-handler, gitCmdAddCommit, cmd_type: %s, cmd_content: %s\n", cmdType, cmdContent)
	var gitResStr []byte

	gitCmd := gitCmdParams{}
	err := json.Unmarshal([]byte(cmdContent), &gitCmd)
	if err != nil {
		return addDto, err
	}

	path := AppRootDir

	// add 操作
	addDto, err1 := c.add(path, &gitCmd)
	fmt.Printf("git-cmd-handler-gitCmdAddCommitHandler, gitAddRes: %+v\n", addDto)
	if err1 != nil {
		gitResStr, _ = json.Marshal(addDto)
		return string(gitResStr), err1
	}

	// commit 操作
	commitDto, err2 := c.commit(path, &gitCmd)
	fmt.Printf("git-cmd-handler-gitCmdAddCommitHandler, gitCommitRes: %+v\n", commitDto)
	if err2 != nil {
		addDto.GitLog += commitDto.GitLog // 将 commit log 拼接到 add log 中
		gitResStr, _ = json.Marshal(addDto)
		return string(gitResStr), err2
	}

	addDto.GitLog += commitDto.GitLog // 将 commit log 拼接到 add log 中
	gitResStr, _ = json.Marshal(addDto)
	return string(gitResStr), nil
}

// 为了保证 engine Marshal 不报错，一定要返回一个结构体
func (c *BaseGitCmd) add(path string, cmd *gitCmdParams) (*gitDTO, error) {
	fmt.Printf("git-cmd-handler-gitCmdAdd, cmd: %+v\n", cmd)
	dto := &gitDTO{}

	startTime := time.Now()
	gitLog, err := c.execGitCmd(path, "git", "add", cmd.File)
	dto.GitLog = gitLog
	if err != nil {
		fmt.Printf("git-cmd-handler-gitCmdAdd, execGitCmd add log:%+v, err: %+v\n", gitLog, err)
		return dto, errors.New(gitLog + err.Error())
	}
	fmt.Printf("git-cmd-handler-gitCmdAdd-success, cmd: %+v, contents: %s, timeSince: %s\n", cmd, gitLog, time.Since(startTime))

	return dto, nil
}

// 为了保证 engine Marshal 不报错，一定要返回一个结构体
func (c *BaseGitCmd) commit(path string, cmd *gitCmdParams) (*gitDTO, error) {
	fmt.Printf("git-cmd-handler-gitCmdCommit, cmd: %+v\n", cmd)
	dto := &gitDTO{}

	startTime := time.Now()
	eCmd := exec.Command("bash", "-c", fmt.Sprintf("source ~/.bashrc && git commit -m '%s'\n", cmd.Message))
	gitLog, err := c.execCmd(path, eCmd)
	dto.GitLog = gitLog
	if err != nil {
		fmt.Printf("git-cmd-handler-gitCmdCommit, execGitCmd commit log:%+v, err: %+v\n", gitLog, err)
		return dto, errors.New(gitLog + err.Error())
	}
	fmt.Printf("git-cmd-handler-gitCmdCommit-success, cmd: %+v, contents: %s, timeSince: %s\n", cmd, gitLog, time.Since(startTime))

	return dto, nil
}

type gitDTO struct {
	GitLog   string `json:"git_log"`   // 用来记录 git 执行 log
	CommitID string `json:"commit_id"` // push 接口用来返回 commit id，其他接口忽略
}

// 在执行 git commit 前设置用户信息
func setGitUserInfo(repoPath string, userName, userEmail string) error {
	cmd := exec.Command("git", "config", "user.name", userName)
	cmd.Dir = repoPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set git user name: %v", err)
	}

	cmd = exec.Command("git", "config", "user.email", userEmail)
	cmd.Dir = repoPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set git user email: %v", err)
	}
	return nil
}

func main() {
	// 使用示例
	err := setGitUserInfo("/home/runner/app", "yunnanwenshan", "clackyai@dao43.com")
	if err != nil {
		fmt.Printf("setGitUserInfo-err: %+v\n", err)
		os.Exit(1)
	}

	cmdGit := BaseGitCmd{}
	for {
		content := "{\"remote_branch\":\"feat/user-module-enhancement\",\"local_branch\":\"feat/user-module-enhancement\",\"user_name\":\"yunnanwenshan\",\"remote_name\":\"origin\",\"message\":\"Fix: Corrected typo in README.md\",\"file\":\".\"}"
		result, err := cmdGit.gitCmdAddCommitHandler("git_add_commit", content)
		if err != nil {
			fmt.Printf("===result: %+v, err: %+v\n", result, err)
		}

		time.Sleep(time.Second * 5)
		fmt.Printf("\n======================sleep: 5\n")
	}
}