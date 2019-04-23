$$.forEach($ => {
    if ($.kind == 'Service') {
        $.metadata.annotations = [{ "cloud66.com/deployed-at": Date.now() }];
    }
});
