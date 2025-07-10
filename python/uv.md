# uv package manager
 ## install
 curl -Ls https://astral.sh/uv/install.sh | bash
 uv --version
 ## basic pip command equivalent in uv
 | pip command                       | uv equivalent                        |
| --------------------------------- | ------------------------------------ |
| `pip install requests`            | `uv pip install requests`            |
| `pip install -r requirements.txt` | `uv pip install -r requirements.txt` |
| `pip freeze > requirements.txt`   | `uv pip freeze > requirements.txt`   |
| `pip uninstall requests`          | `uv pip uninstall requests`          |
| `pip list`                        | `uv pip list`                        |
| `pip show requests`               | `uv pip show requests`               |


## managing virtual environment
uv venv .venv
source .venv/bin/activate
