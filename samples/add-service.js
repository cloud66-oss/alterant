$$.forEach($ => {
    var namespace = $.metadata.name
    deployment = {
        apiVersion: "extensions/v1beta1",
        kind: "Deployment",
        metadata: [
            { namespace: namespace },
            { name: "web" }
        ],
        spec:
            { template:
                { spec:
                    { containers: [{ "image": "app_image:latest", "name": "my-pod" }] }
                }
            }
        }

    $$.push(deployment);
});
