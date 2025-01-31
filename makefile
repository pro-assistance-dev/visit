pre_commit:
	./cmd/git-hooks/pre-commit	

#GIT#
#####
#####

git_push: 
	./cmd/git/git_push.sh "$m"

git_commit:
	./cmd/git/git_commit.sh "$m"

git_merge: git_push
	git checkout develop
	git pull
	git merge @{-1}
	git push

# example: make git_feature n=1
git_feature:
	git flow feature start $n

