## REST

### add an app to the store

- `http://{{rubix_apps_ip}}:{{rubix_apps_port}}/api/stores`
- `POST`

#### allowable_products

if an app is added to the store and allowable_products is meant to be used so the user can install an app on the
incorrect target device

```
{
    "name": "flow-framework",
    "allowable_products": [
        "RubixCompute",
        "AllLinux"
    ],
    "Port": 1660,
    "app_type_name": "Go",
    "repo": "flow-framework",
    "service_name": "nubeio-flow-framework",
    "RubixRootPath": "/data",
    "apps_path": "/data/rubix-apps/installed",
    "app_path": "/data/flow-framework",
    "download_path": "/home/aidan/downloads",
    "asset_zip_name": "",
    "owner": "NubeIO",
    "run_as_user": "root",
    "data_dir": "",
    "config_dir": "",
    "config_file_name": "",
    "description": "flow-framework",
    "service_working_directory": "/data/rubix-apps/installed/flow-framework",
    "service_exec_start": "/data/rubix-apps/installed/flow-framework/app-amd64 -p 1660 -g /data/flow-framework -d data -prod",
    "product_type": "status",
    "arch": ""
}

```

### install an app

- `http://{{rubix_apps_ip}}:{{rubix_apps_port}}/api/apps`
- `POST`

```
{
    "app_name": "flow-framework",
    "version":"latest",
    "token":""
}
```

#### get progress

- `http://{{rubix_apps_ip}}:{{rubix_apps_port}}/api/apps/progress/install`
- `POST`

```
{
    "app_name": "flow-framework"
}
```

### uninstall an app

- `http://{{rubix_apps_ip}}:{{rubix_apps_port}}/api/apps`
- `DELETE`

```
{
    "app_name": "flow-framework"
}
```

#### get progress

- `http://{{rubix_apps_ip}}:{{rubix_apps_port}}/api/apps/progress/uninstall`
- `POST`

```
{
    "app_name": "flow-framework"
}
```