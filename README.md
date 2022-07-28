# EdRank Backend


- [API DOCS](https://github.com/edrank/edrank_backend/blob/master/docs/API_DOCS.md)
- [DEV API DOCS](https://github.com/edrank/edrank_backend/blob/master/docs/DEV_API_DOCS.md)

## Setting up locally

- Clone the repo
- Install Docker and Docker Compose from [here](https://docs.docker.com/compose/install/)
- Check installation - `docker ps`, `docker images`, `docker-compose -v`
- Move to root directory of the project
- Spin up the backend
    - On windows
        - `scripts/spin.bat` (without rebuilding)
        - `scripts/spin.bat --build` (rebuild image)
    - On Linux
        - `bash spin.sh`
- To stop the backend
    - On windows
        - `scripts/down.bat`
    - On Linux
        - `bash down.sh`

### Version Control 
- Take a pull from master branch - `git pull`
- Run `scripts/spin.bat --build`
- **If there are no new changes, no need to rebuild image** 
    - Just run `scripts/spin.bat`