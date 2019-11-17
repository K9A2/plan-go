# plan-go

`plan-go` is a simple and stupid plan management utility tool.

## Usage

```txt
Add a new plan
$ plan add `future plan`
// $ major plan 'future plan' (11ff22) added, 

Add a new child plan as a child of parent plan 11ff22
$ plan add 'future child-plan' -p 11ff22
// $ child plan 'future child-plan' (11ff33) appended to parent plan `future plan` (11ff22)

Mark a plan and all its children as done
$ plan done 11ff22
// $ plan `future plan` (11ff22) marked as done

List all plan
$ plan status
// $ 1. major plan: 11ff22, date: Sat Nov 16 15:43:47 2019 +0800
// $   future plan
// $   1.1 child plan: 11ff33, date: Sat Nov 16 15:43:47 2019 +0800
// $     future child-plan
// $ ...

Reopen a plan and all its children as undone
$ plan reopen 11ff22
// $ plan `future plan` (11ff22) reopened

Delete a plan and all its children
$ plan delete 11ff22
// $ plan `future plan` (11ff22) deleted

Retitle a plan
$ plan retitle 11ff22 `a new title`
// $ plan 11ff22 retitled as `a new title`
```

# Build

Run `make` to execute makefile, and built file will be placed in `build/`. Run 
as root to gain permission for install/uninstall operations.

```text
Install built files to /usr/local/plan, and create link at /usr/bin/plan
# make install 

Uninstall program
# make uninstall

Clean built destination
# make clean
```
