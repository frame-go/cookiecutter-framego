# System Design

## API Design

```
{{ cookiecutter.__service_name_title }} {
    CreateObject [post /{{ cookiecutter.service_name }}/v1/object/create]: Create object
    GetObject [get /{{ cookiecutter.service_name }}/v1/object/get]: Get object by name
    UpdateObject [post /{{ cookiecutter.service_name }}/v1/object/update]: Update object
    DeleteObject [post /{{ cookiecutter.service_name }}/v1/object/delete]: Delete object
}
```

## DB Design

```
@{{ cookiecutter.service_name }}_db

object_tab
uint64 id(auto)
varchar[50] name
varchar[1000] data(charset=utf8mb4)
uint32 create_time
uint32 update_time
primary key(id)
unique key idx_name(name)
```
