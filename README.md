# turgo
> Command line bookmarker

Turgo uses **[Turso](https://turso.tech)** to host your data

## Install
<details>
<summary>Manually</summary>
<br/>
  
```bash

git clone https://github.com/musaubrian/turgo
cd turgo

go build .

# run it

turgo
```
</details>

<details>
<summary>go install </summary>
<br/>
  
```bash

go install github.com/musaubrian/turgo@latest
# run it
turgo
```
</details>

### Setup
> **Note**
> 
> You need to have **[turso-cli](https://docs.turso.tech/tutorials/get-started-turso-cli)** installed to get the token and db URL

To configure the environment variables, simply run `turgo` and it'll prompt you
to enter the required details
