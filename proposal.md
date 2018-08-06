# Administering large scale PostgreSQL installations on SAP MultiCloud platform

Postgresql-as-a-Service is provided on large scale on SAP MultiCloud platform. The platform supports cloud provider(s) like aws, azure, gcp and openstack. Together on all infrastructures we provide approximately 5000 clusters of postgresql service.

### Postgres cluster is robust and intelligent enough to remain up and running

In a cluster we have two postgresql nodes running. One node runs as a primary and other node runs as a secondary (also called as a hot standby). Secondary node is replicating primary with replication set to asynchronous mode. Primary node failure is mitigated by promoting secondary node to primary. So cluster is up and running always. To detect the the failure pgpool is used which continusly checks the heartbeat of postgresql process. We have pgpool running in three nodes inorder to form consensus. Refer below figure for more understanding.

[![N|Solid](https://github.com/dbossap/dbos-performance/blob/master/clusterSetup.png?raw=true)](https://nodesource.com/products/nsolid)

### Disaster Recovery situation is handled seamlessly with the help of Backup and Restore (B&R) and High Availability (HA) features

Backup and Restore feature helps user to take online backup of running cluster on MultiCloud platform. The backup approach will differ for respective cloud providers. For e.g. on OpenStack we use 'tarball' approach, for other cloud providers (aws, azure and gcp) we use 'snapshot' based approach. Incase of failure/disaster situation this backup can be restored to avoid any data lose. Also at any point of time user can move the cluster to previous state(data) by restoring appropriate backup.

High availability (HA) feature helps to mitigate the failure situation. HA detects primary node failure and promote secondary node to primary. Split brain is the well know and dangerous problem that could occur in this architecture (Primary-secondary). We have avoided this problem using STONITH operation. Here we kill the failed node after promoting the secondary node to primary.

Client/Customer always connect to single endpoint which always points to primary node. Incase of failure situation, HA makes sure that single endpoint provided to client will always point to primary node with minimal downtime.

### Integration with [bosh] and [service-fabrik] makes Postgresql-as-a-Service a scalable and independent component on MultiCloud platform

[Bosh] is an important component of SAP MultiCloud platform. It is an open source project that offers a tool chain for release engineering, deployment & life-cycle management of large scale distributed services. All the postgresql clusters are deployed using bosh on MultiCloud platform. Bosh also helps in accessing and maintaining the cluster up and running. All the bosh life-cycle operations are triggered through [service-fabrik].

[Service-fabrik] is an open source component of SAP MultiCloud platform which acts as a broker between customers/clients and Postgresql-as-a-service. All the service operations are triggered through [service-fabrik]. Some operations are scheduled for e.g. scheduled backup, cluster security updates etc. Operations such as create cluster, delete cluster, update cluster, upgrade cluster will be triggered by end user customers/clients through [service-fabrik].

### Updating Postgresql-as-a-Service with new features and OS security patches is fairly easy with rolling updates (bi-weekly updates) of clusters

SAP MultiCloud platform supports rolling updates of clusters. Every postgresql cluster is updated bi-weekly with new features or bug fixes (if any). This keeps us improving. 
OS security patches will be applied on all clusters once in a month. With this update we get rid of security vulnerability (if any).

### SAP MultiCloud platform monitors health of every postgresql cluster

Monitoring component captures various metrics of the postgresql cluster which helps to monitor the health of the each cluster. We capture some of the core metrics like memory usage, cpu usage, disk usage etc. Also metrics such as number active connections, bulk data read, replication status, availability, failover status (if any), backup status gives detail cluster information to identify any failure.

All the metrics are captured by an agent running inside postgresql node and sent to riemann. Riemann pushes these metrics to influxdb where all the metrics will be stored. Grafana is used to show all the metrics on its dashboard. Grafana fetches these metrics from influxdb and displays it. Grafana gives user flexibility to select cluster, select date and time range etc. Monitoring is one of the important component SAP MultiCloud platform. Refer below figure

[![N|Solid](https://github.com/dbossap/dbos-performance/blob/master/grafana.png?raw=true)](https://nodesource.com/products/nsolid)


### Altering module of SAP MultiCloud platform quickly identifies and notifies the failures

As a part of alerting module we have configured rules in riemann on metrics like availability, backup, disk usage etc. When these metrics does not adhere to these riemann rules an alert will be raised.

### Debugging is independent of postgrsql cluster availabilty

All the logs generated out of service will be pushed to ELK stack so that user can access these logs incase of any failure. So debugging is independent of service instance's availability. Refer below figure

[![N|Solid](https://github.com/dbossap/dbos-performance/blob/master/kibana.png?raw=true)](https://nodesource.com/products/nsolid)

### SAP MultiCloud platform supports variety of plans with major version upgrade

Currently we support both 9.4 and 9.6 version's  of postgresql with different t-shirt sizes (plans). As per requirement user has the flexibility to choose the plan. Plan provided will be like v9.4-small, v9.6-small, v9.6-large etc. Each plan is unique combination of disk, cpu, memory etc. 
We also support major version upgrade of postgresql service. This feature easily allows user to upgrade its postgresql instance to next higher version.

# New Features!

  - Import a HTML file and watch it magically convert to Markdown
  - Drag and drop images (requires your Dropbox account be linked)


You can also:
  - Import and save files from GitHub, Dropbox, Google Drive and One Drive
  - Drag and drop markdown and HTML files into Dillinger
  - Export documents as Markdown, HTML and PDF

Markdown is a lightweight markup language based on the formatting conventions that people naturally use in email.  As [John Gruber] writes on the [Markdown site][df1]

> The overriding design goal for Markdown's
> formatting syntax is to make it as readable
> as possible. The idea is that a
> Markdown-formatted document should be
> publishable as-is, as plain text, without
> looking like it's been marked up with tags
> or formatting instructions.

This text you see here is *actually* written in Markdown! To get a feel for Markdown's syntax, type some text into the left window and watch the results in the right.

### Tech

Dillinger uses a number of open source projects to work properly:

* [AngularJS] - HTML enhanced for web apps!
* [Ace Editor] - awesome web-based text editor
* [markdown-it] - Markdown parser done right. Fast and easy to extend.
* [Twitter Bootstrap] - great UI boilerplate for modern web apps
* [node.js] - evented I/O for the backend
* [Express] - fast node.js network app framework [@tjholowaychuk]
* [Gulp] - the streaming build system
* [Breakdance](http://breakdance.io) - HTML to Markdown converter
* [jQuery] - duh

And of course Dillinger itself is open source with a [public repository][dill]
 on GitHub.

### Installation

Dillinger requires [Node.js](https://nodejs.org/) v4+ to run.

Install the dependencies and devDependencies and start the server.

```sh
$ cd dillinger
$ npm install -d
$ node app
```

For production environments...

```sh
$ npm install --production
$ NODE_ENV=production node app
```

### Plugins

Dillinger is currently extended with the following plugins. Instructions on how to use them in your own application are linked below.

| Plugin | README |
| ------ | ------ |
| Dropbox | [plugins/dropbox/README.md][PlDb] |
| Github | [plugins/github/README.md][PlGh] |
| Google Drive | [plugins/googledrive/README.md][PlGd] |
| OneDrive | [plugins/onedrive/README.md][PlOd] |
| Medium | [plugins/medium/README.md][PlMe] |
| Google Analytics | [plugins/googleanalytics/README.md][PlGa] |


### Development

Want to contribute? Great!

Dillinger uses Gulp + Webpack for fast developing.
Make a change in your file and instantanously see your updates!

Open your favorite Terminal and run these commands.

First Tab:
```sh
$ node app
```

Second Tab:
```sh
$ gulp watch
```

(optional) Third:
```sh
$ karma test
```
#### Building for source
For production release:
```sh
$ gulp build --prod
```
Generating pre-built zip archives for distribution:
```sh
$ gulp build dist --prod
```
### Docker
Dillinger is very easy to install and deploy in a Docker container.

By default, the Docker will expose port 8080, so change this within the Dockerfile if necessary. When ready, simply use the Dockerfile to build the image.

```sh
cd dillinger
docker build -t joemccann/dillinger:${package.json.version}
```
This will create the dillinger image and pull in the necessary dependencies. Be sure to swap out `${package.json.version}` with the actual version of Dillinger.

Once done, run the Docker image and map the port to whatever you wish on your host. In this example, we simply map port 8000 of the host to port 8080 of the Docker (or whatever port was exposed in the Dockerfile):

```sh
docker run -d -p 8000:8080 --restart="always" <youruser>/dillinger:${package.json.version}
```

Verify the deployment by navigating to your server address in your preferred browser.

```sh
127.0.0.1:8000
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT


**Free Software, Hell Yeah!**

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)


   [bosh]: <https://github.com/joemccann/dillinger>
   [pgpool]: <https://github.com/joemccann/dillinger.git>
   [grafana]: <https://grafana.com/>
   [riemann]: <http://riemann.io/>
   [influxdb]: <https://www.influxdata.com/time-series-platform/influxdb/>
   [service-fabrik]: <https://github.com/cloudfoundry-incubator/service-fabrik-broker>

