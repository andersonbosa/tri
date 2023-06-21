<section align="center">

  <img src="docs/banner.svg" title="Project banner" alt="Project banner" />

  <br>
  <br>

  <!-- badges -->

  <p>
    <a href="#about">About</a> •
    <a href="#technologies">Technologies</a> •
    <a href="#getting-started">Getting Started</a> •
    <a href="#contribution">Contribution</a> •
    <a href="#license">License</a>
  </p>
</section>

---

<h2 id="about">💬 About</h2>

Building a Todo application in Go using Cobra

> Building an Awesome CLI App in Go – OSCON 2017
> A workshop written and delivered by Steve Francia and Ashley McNamara at OSCON 2017 outlining the techniques, principles, and libraries you need to create user-friendly command-line interfaces and command suites before walking you through building your own app. Along the way, you’ll cover everything from how to design and build commands to working with and parsing flags, config files and remote config systems, and how to work with environment variables and 12-factor apps. By the end of the workshop, you’ll have a working knowledge of Go and your very own functioning CLI app.
> 
> This workshop covers:
> 
> - CLI application design
> - Introduction to the Go programming language
> - Introduction to the Cobra CLI framework (used by Kubernetes, Docker, Hugo, etc)
> 
> Source: https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/

<h2 id="technologies"> 🛠️ Technologies</h2>

- [Go language](https://go.dev/)
- [Cobra](https://github.com/spf13/cobra)
- [Viper](https://github.com/spf13/viper)

<h2 id="getting-started"> 🟢 Getting Started </h2>

### Installation

```bash
# download latest version
wget -L https://github.com/andersonbosa/tri/raw/main/tri/tri -O $HOME/.local/bin/tri

# give executino permission
chmod +x $HOME/.local/bin/tri

# validate instalation
tri --help
```

### Usage

#### Help message

```
Tri will help you get more done in less time.
It's designed to be as simple as possible to help
you accomplish your goals.

Usage:
  tri [command]

Available Commands:
  add         Add a new todo
  done        Mark item as Done
  help        Help about any command
  list        List all todos

Flags:
      --config string     config file (default is $HOME/.tri.yml)
      --datafile string   data file to store todos (default "/home/t4inha/.tri_todo.json")
  -h, --help              help for tri

Use "tri [command] --help" for more information about a command.
```

<h2>🚀 Deploy</h2>

* Distributed here, in [Github](https://github.com/andersonbosa/tri/releases)

<h2 id="contribution">🤝 Contribution</h2>

<p>
  This project is for study purposes too, so please send me a message telling me what you are doing and why you are doing it, teach me what you know. All kinds of contributions are very welcome and appreciated!
</p>


<h2 id="license"> 📝 License</h2>

This project is under the Unlicense.


---

<h4>  
  <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/andersonbosa/tri?style=social">
  | Did you like the repository? Give it a star! 😁
</h4>
