![img.png](img.png)

# DuckVault
___
## A CLI tool to convert your beloved data into K8 Secrets

### Don't worry! I'm gonna take care of you


### Features
- Convert multiple data at once into secrets
- Provides base64 encoding for secrets
- Currently working from ENV files
- Converting into YAML files
- Define Namespaces


## Usage

DuckVault provides the ability to create multi secrets from env files based on the content of that file. With the generate command you are able to do that. On default it will search for a env file in your current directory
```
duckvault generate
```

If you have to specify a path there is the option to use the filePath flag to set the path to your env file.
```
duckvault generate --f /path/to/your/.env/file
```

Kubernetes provides you with the ability to order your data into namespaces. With the ns flag you can set it manually. On default it is set to default
```
duckvault generate --ns namespace
```



## License
This project is licensed under the [MIT License](LICENSE).