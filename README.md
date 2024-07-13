# mark
A re-implementation of bookmark using bubble tea!

## Requirements

- `go` installed

## Install

```bash
git clone https://github.com/josiahdenton/mark.git
cd mark
go install .
```

You may have to add `$HOME/go/bin` to your path if you have not yet.
Once complete, you can run `mark` from anywhere.

## Usage

`mark` uses a sqlite DB in `~/.mark/mark.db` to persist bookmarks. You can (a)dd, (e)dit, or (d)elete 
any of the marks. If you accidentally delete something, hit (u)ndo to add the mark back.

##### List (Book)marks 
<img width="675" alt="image" src="https://github.com/user-attachments/assets/d217c547-3766-45cc-b8ab-96ffc6cbe50e">

##### Search through marks
<img width="679" alt="image" src="https://github.com/user-attachments/assets/ca1c2bd0-224f-458b-b421-e62461ac71d9">

##### Add/Edit marks
<img width="683" alt="image" src="https://github.com/user-attachments/assets/71ddf8d7-32fd-4cc1-a97c-7b21cc610dcb">

- ran these in zed terminal window


## Progress

- [ ] fix all unit tests
- [ ] max size on titles set (to avoid NLs!)
- [ ] add help menu

## CONTRIBUTING

Opening issues welcome.
