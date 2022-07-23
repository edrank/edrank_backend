# EdRank Backend


[API DOCS](https://github.com/edrank/edrank_backend/blob/master/docs/API_DOCS.md)

## Setting up locally

- Clone the repo
- Install go [here](https://go.dev/dl/)
- Check go installation - `go version`
- Install AWS CLI [here](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
- Configure AWS CLI by typing - `aws configure`
- Fill in the prompt values from the `./config/aws.json` file
- Run `cat ~/.aws/credentials` to verify the values
- Move to root directory of the project
- On windows
    - `spin.bat`
- On Linux
    - `bash spin.sh`