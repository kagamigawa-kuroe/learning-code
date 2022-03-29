# GIT avance

### 记录一些git的高级用法以及在工作学习中遇到的问题（持续更新）

1. > Q: 在第一次提交项目的时候，不慎从master分支(而非自己fork出的远程仓库分支中)fetch了内容，导致本地出现了很多不必要的分支，同时，也与master分支中的内容有出入。

   > A: 首先 删除全部的分支
   >
   > `git branch | grep -v "master" | xarg git branch -D`
   >
   > ​    然后下载远程master分支中的内容到本地分支
   >
   > ` git checkout -b nxdmaster /remotes/nexedi/master`
   >
   > ​    最后将自己的分支合并到master上
   >
   > `git rebase nexdmaster matomo`
   >
   > ​    这样rebase之后 发现无法合并 因为本地有许多除了matomo分支之外的文件和远程仓库master中的内容无法对上，所以我们需要根据提示逐条修改。
   >
   > `git status` 查看具体有哪些文件在rebase过程中产生冲突
   >
   > `git log ../wendelin/software.cfg` 查看产生冲突的文件log
   >
   > `git checkout XXX ../../stack/monitor/buildout.cfg` 将产生冲突的文件恢复到具体的某个版本
   >
   > `git rebase --continue` 修改完成后继续rebase
   >
   > ​    当然，这是一种较为麻烦的做法，尤其是当产生的冲突较多的时候，要修改的文件会很多。
   >
   > ​    因此，我们也可以采取另一种做法，如下
   >
   > `git rebase --abort` 我们先结束rebase的过程
   >
   > `git checkout -b matomo-backup` 保险起见 先备份
   >
   > `git reset --hard bd16bf2173c9171ce96b5e507c92aaddfdd34f4c` 将分支内容恢复到最初始的版本，也就是新增内容没有加入的版本
   >
   > `git cherrypick XXXX` 取回前面几次提交的相关内容，注意，这里取回的是变化的内容，也就是将提交中发生的变化移植到现在的版本上
   >
   > 这样 我们就解决了问题
   >
   > `git push -f` 最后 我们提交内容 解决问题

2. > Q. 在提交的过程中，有些多无意义的提交或者内容相似的提交，想要整合一下让整个提交的目录看起来更为简洁明了

   > A. `git rebase -i XXX`
   >
   > ​    XXX为你想要合并的提交的最早一次提交的前一次
   >
   > ​    然后会进入到一个文件
   >
   > ​    将你想合并的文件的开头` pick` 改为 `squash`, 确认后即可，他将会和上一次提交合并。
   >
   >    同时还有一系列别的用法。
   >
   >    比如将`pick`改成`reword`可以更改commit的评论信息等等        

3. > Q. 在提交完整个项目后，发现存在许多错误的提交，想要从零开始重新整理提交到远程仓库。
   >
   > A. 找到最早的一次提交的前一次提交，使用reset命令来恢复。
   >
   > ​     一开始，我直接食用`git reset --hard xxx`来实现恢复，然后再使用 `git rm --cache filename ` 来依次取消提交，这样做比较麻烦。
   >
   > ​     后来发现如果想将暂存区的信息一并恢复到某个版本，可以使用reset的mixed参数，这也是reset的默认参数，可以升落
   >
   > ​     `git reset --mixed xxx`
   >
   > ​     这样之后，再重新整理add并且commit。