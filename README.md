# Quithub
### Lightweight and simple local version control for when you're too lazy to use GitHub 🤷‍♀️🤷‍♂️
#### Let's be real, how many commands do you know besides: 
#### 1. `add .`
#### 2. `commit -m "vague, unhelpful description here to piss off your team"`
#### 3. `push`?
#### Okay, maybe I just never learned how to properly use `git` from the command line 😭😝

*P.S. I was pretty close to naming this project 💩Hub.*

___

### Installation
To start using `quithub`, install Go and run `go get`:

`$ go get github.com/AndrewYinLi/QuitHub`

___

### Usage
**Begin by navigating to within the directory to backup:**

`cd <directory path>`

**Back up current working directory under a new name:**

`quithub commit <name of commit>`

**Revert current working directory to a past commit:**

`quithub revert <name of commit>`

**Delete a commit from history:**

`quithub delete <name of commit>`

**Print commit history:**

`quithub history`
