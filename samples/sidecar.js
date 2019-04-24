$$.forEach(function($) {
    // adds a CloudProxy sidecar to the deployment
    if ($.kind === 'Deployment') {
        var containers = $.spec.template.spec.containers;
        if (containers.length === 1) {
            sidecar = JSON.parse(JsonReader("samples/sidecar.json"));

            var sidecarImage = new DockerImage(sidecar.image);
            var containerImage = new DockerImage(containers[0].image);

            sidecarImage.tag = containerImage.tag;
            sidecar.image = sidecarImage.address();
            containers.push(sidecar);
        }
    }
});
