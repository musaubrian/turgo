# turgo
> Command line bookmarker

Turgo uses **[turso]**(https://turso.tech) to host your data

## Install
<details>
<summary>Manually</summary>


```bash
git clone https://github.com/musaubrian/turgo
cd turgo

go build .

# run it

turgo
```
</details>

<details>
<summary>Using `go install`</summary>

```bash
go install github.com/musaubrian/turgo@latest
# run it
turgo
```
</details>

### Setup
> **Note**
> 
> You need to have **turso-cli** installed to get the token and db URL

To configure the environment variables, simply run `turgo` and it'll prompt you
to enter the required details

You can get them **[here]**(https://docs.turso.tech/tutorials/get-started-turso-cli)
