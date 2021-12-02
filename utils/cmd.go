package main

import (
    "context"
    "os/exec"
    "syscall"
)

func RunCmd(ctx context.Context, cmd *exec.Cmd) error {
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Setpgid: true,
    }

    if err := cmd.Start(); err != nil {
        return err
    }

    errCh := make(chan error, 1)
    go func() {
        errCh <- cmd.Wait()
    }()

    done := ctx.Done()
    for {
        select {
        case <-done:
            done = nil
            pid := cmd.Process.Pid
            if err := syscall.Kill(-1*pid, syscall.SIGKILL); err != nil {
                return err
            }
        case err := <-errCh:
            if done == nil {
                return ctx.Err()
            } else {
                return err
            }
        }
    }
}
