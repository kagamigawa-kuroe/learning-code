# GIT remote

> 从仓库克隆某个项目
>
> `git clone url`

> 添加远程仓库
>
> ```git
> git remote add name url
> ```
>
> 移除远程仓库
>
> ```g
> git remote rm name
> ```
>
> 从远程仓库下载新分支与数据
>
> ```
> git fetch
> ```
>
> 从远端仓库提取数据并尝试合并到当前分支：
>
> ```
> git pull = git fetch + git merge
> git pull --rebase = git fetch + git rebase
> ```

> `git push <仓库名> <本地分支名>:<远程分支名>` 将本地分支关联到远程分支
>
> `git branch -r -d origin/远程分支名   ` 删除远程分支
>
> `git push origin :远程分支名`  删除远程分支
>
> `git checkout --track origin/branch_name`  同时在本地和远程创建branch_name分支

> 如果本地新建了一个分支 branch_name，但是在远程没有
>
> `git push --set-upstream origin new_branch_name`  可以创建
>
> 实现在本地删除远程已经不存在的分支
>
> `git fetch --prune`

