
Go 依赖管理的三个阶段 GOPATH, GOVENDOR, go mod

Go code is grouped into packages, and packages are grouped into modules.

### Create a module

Start your module using the go mod init command.
Run the go mod init command, giving it your module path -- here, use example.com/greetings. If you publish a module, this must be a path from which your module can be downloaded by Go tools. That would be your code's repository.

For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the example.com/greetings code on your local file system.

