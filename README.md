# Alterant

Alterant is a tool that transforms configuration files based on custom scripts.  It reads configuration files in yaml or json and modifies them based on your scripts.


## Welcome to Alterant

Alterant is a Descriptive Configuration Modifier. It reads configuration files in <code>yaml</code> or <code>json</code> and modifies them based on your scripts. You can think of it as an elegant and understandable equivalent of XSLT but for YAML!

### Why do I need Alterant?

While Alterant can be used with any YAML file, it is most useful if you use it with Kubernetes configuration files. Kubernetes configuration files describe how infrastructure resources should be deployed and configured on a cluster. However, a lot of times we find ourselves applying the same modifications to many of them.

#### Here is an example: Make sure all pods have log collection

If you want to collect your pod's logs and send them to a log management facility, you might want to use the [Sidecar pattern](https://kubernetes.io/docs/concepts/cluster-administration/logging/) to include a log collector in every pod that's deployed. To do this you will need to remember to add the log collection sidecar to all pods and configure them correctly. If your log collector configuration needs a change, you'd need to go around and make that change everywhere. That's not good for obvious reasons:

1. Adding the sidecar is cumbersome and prone to errors
2. Any change needs to be applied everywhere
3. You can't apply them as part of your CI/CD based on environment unless you keep different configuration files for each environment.

**Deploying log collection without Alterant**

*Your deployment without a sidecar:*
```yaml
apiVersion: v1
kind: Pod
metadata:
	name: counter
spec:
	containers:
	- name: count
		image: my_app
		args: ['--log-folder', '/var/log/app.log']
		volumeMounts:
		- name: varlog
			mountPath: /var/log
	volumes:
	- name: varlog
		emptyDir: {}
```

*Your deployment with a sidecar:*
```yaml
apiVersion: v1
kind: Pod
metadata:
	name: counter
spec:
	containers:
	- name: count
		image: my_app
		args: ['--log-folder', '/var/log/app.log']
		volumeMounts:
		- name: varlog
			mountPath: /var/log
	- name: log-collector-sidecar
		image: my_log_collector
		args: ['--collect-from', '/var/log/app.log']
		volumeMounts:
		- name: varlog
			mountPath: /var/log
	volumes:
	- name: varlog
		emptyDir: {}
```

As you can see, the following piece needs to be added to every Deployment configuration in your app:

```yaml
	- name: log-collector-sidecar
		image: my_log_collector
		args: ['--collect-from', '/var/log/app.log']
		volumeMounts:
		- name: varlog
			mountPath: /var/log
```

**Deploying log collection with Alterant**

With Alterant, we can write a simple Javascript file to add the sidecar to any input Deployment resource:

```javascript
sidecar = YamlReader("sidecar.yaml")
$$.forEach(functin($) {
	$.spec.template.spec.containers.push(sidecar)
});
```

I can then put the sidecar configuration in a file called `sidecar.yaml`:

```yaml
	- name: log-collector-sidecar
		image: my_log_collector
		args: ['--collect-from', '/var/log/app.log']
		volumeMounts:
		- name: varlog
			mountPath: /var/log
```

Now all I have to do is to run Alterant:

```bash
alterant modify --in deployment.yml --modifier sidecar.js
```

#### ...but what about Admission Controllers?

If you work with Kubernetes, you might have heard about [Admission Controllers](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/). Using Admission Controllers you can add a rule to your cluster to make changes to the input requests, like creation of a Deployment, so any Deployment will get a sidecar added to it automagically. While that's a way of doing the same thing and you're more than welcome to do it, we prefer keeping simple things simple and more importantly taking the magic out of automagic. Admission Controllers are powerful and can do a lot. They can also confuse us since they add extra things to our configuration without our knowledge, unless we know that they are present. Also, they need to be added to your cluster and that's not always possible or desired.

## How can I get started?

See the Getting started section in the wiki section of this repository.

## What's Alterant NOT for?

Being able to change configuration files with a script might tempt us to do it more often than we should! Obviously it is up to you how you use Alterant, but we have some suggestions as when to use and not use use Alterant.

1. Use Alterant as an automatic step to add or modify generic and repetitive configuration parts to your configuration files.
2. Use Alterant to make sure your configuration files adhere to your best practices or policies. If validation of configuration files is your main purpose for using Alterant, you can use [check out Copper](/copper/index.html).
3. Do not use Alterant to change a configuration file to another just because you don't want to create a duplicate. An example is when your Kubernetes configuration files are slightly different between 2 clusters. You should use placeholders, templates or [Cloud 66 Skycap](https://cloud66.com/containers/skycap) for that!
4. Do not use Alterant to obfuscate changes to configuration files as an automatic step. Alterant applies changes to the configuration files before they are applied to the application (ie. your Kubernetes cluster) exactly because it wants to make those changes transparent. Don't use Alterant as a magic step!


## What else can it do?

A lot! Alterant safely runs your Javascript code to manipulate the input file. This gives you a lot of power and flexibility. It also provides some shortcuts and helpers to deal specifically with Docker and Kubernetes configuration files.

Alterant can be used in a pipeline or generate **diff** files just so you can see what would happen it it ran on the file without making changes. It can also convert your configuration files between YAML and JSON.

## Show me some examples

Sure! Checkout the wiki section.
