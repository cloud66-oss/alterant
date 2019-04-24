istioSidecar = JSON.parse(YamlReader("samples/istio.yml"));

// go through all items
$$.forEach(function($) {
    sidecar = istioSidecar;

    // if you find a service, then look for it's deployment
    if ($.kind === "Service") {
        // we have a service. look for it's deployment
        selectors = $.spec.selector;
        deployment = findDeploymentForService($$, selectors);

        // find the service name and add it as the last arg of istio container config
        var name = $.metadata.name;
        sidecar.args.push(name);

        // add the side car to the deployment
        var containers = deployment.spec.template.spec.containers;
        if (containers.length === 1) {
            containers.push(sidecar);
        }

        $$.replace(deployment);
    }
});
