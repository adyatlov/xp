Bun Explorer
============

Developing
----------

1. Start server:

    ```bash
    $ go build && ./bunxp <path to bundle>
    ```
   
    The best way to test the GraphQL API is https://github.com/hasura/graphqurl:
    ```bash
    $ gq http://localhost:7777/graphql -i 
    ```
   
2. Start client:

    ```bash
    $ cd client
    $ yarn install # only the first time
    $ yarn run relay --watch
    $ yarn start # in a separate terminal
    ```

    You don't need to relaunch the `npm start` command after you changed client files,
    they will be reloaded automatically.
    
3. Open `http://localhost:3000`

Build project
-------------

1. Install [Packr v2](https://github.com/gobuffalo/packr/tree/master/v2):

    ```bash
    $ go get -u github.com/gobuffalo/packr/v2/packr2
    ```

3. Build client:

    ```bash
    $ cd client
    $ npm run build 
    ```

4. Add client files to server binaries:

    ```bash
    packr2 
    ```
   
5. Build client-server bundle:

    ```bash
   go build
    ```
   
6. Clean-up:

    ```bash
    packr2 clean
    ```
   
Object Types
------------

- Cluster
- Agent
- Framework
- Task
- Marathon App
- Component
- API Endpoint Output
- Command Output
- Log
- File
- Role

Metric Types
------------

- Integer
- Real
- Percentage
- Version
- Timestamp
- IP Address

