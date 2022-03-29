##                                         GIT base

> 1. 创建新的git目录
>
> ​        `git init`
>
> ​        path目录将会被视为git目录，自动生成.git隐藏文件夹

> 2. 将修改提交到暂存区( Stage/index )
>
>    `git add XXX XXX`
>
>    `git add ./*`
>    
>    在正式提交前 可以多次add不同的文件

> 3. 提交
>
>    `git commit -m "xxxx"`
>    
>    xxxx是你的提交注解

> 4. 查看历史记录和当前状态
>
>    其中记录了每一次提交的号码, 以及提交内容
>    
>    `git log/git status`

> 5. 回溯
>
>    `git reset --hard HEAD^` (回溯到上一个版本)
>
>    `git reset --hard HEAD^^` (回溯到上上个版本)
>
>    `git reset --hard HEAD~4`(回溯到4个提交前的版本)
>    
>    `git reset --hard XXXXXX` (根据log上的更改编号来指定版本）

> 6. `git checkout --file`
>
>    可以丢弃工作区的修改
>
>    用暂存区/版本库里的版本替换工作区的版本，无论工作区是修改还是删除，都可以“一键还原”。
>
>    
>
>    `git reset [--soft | --mixed | --hard] [HEAD] `
>    
>    mixed为默认参数 可以把暂存区的内容用某个版本替换 工作区内容保持不变
>    
>    soft 用于退回到某个版本 工作区内容保持不变
>    
>    hard 退回到某个版本 工作区内容也会被更改

> 7. `git rm`
>
>    从版本库中删除一个文件

> 8. 分支
>
>    `git branch <name>` 创建分支
>
>    `git branch -a` 查看所有分支(包括远程分支)
>
>    `git checkout <name>/git switch <name>` 切换分支
>
>    `git merge <name>` 合并某分支到当前分支
>
>    `git rebase <name>`合并当前分支到某分支上
>
>    `git branch -d <name>` 删除分支
>
>    `git checkout -b <name>`创建并切换到该分支
