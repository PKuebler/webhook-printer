function NewClient() {
    let conn = null
    const listener = {}

    function connect(host) {
        conn = new WebSocket("ws://" + host + "/ws");

        conn.onclose = function (evt) {
            trigger({
                name: "_close",
                event: ""
            })

            setTimeout(function() {
                connect(host)
            }, 1000)
        };

        conn.onerror = function(err) {
            conn.close()
        }

        conn.onopen = function(evt) {
            trigger({
                name: "_open",
                event: ""
            })
        }

        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                if (messages[i] == "") {
                    continue;
                }
                msg = JSON.parse(messages[i]);
                console.log("trigger", msg)

                trigger(msg)
            }
        }
    }

    function trigger(msg) {
        let eventUpdate = ""
        if (msg.event && msg.event != "") {
            eventUpdate = JSON.parse(msg.event);
        }
        if (listener.hasOwnProperty(msg.name)) {
            listener[msg.name].forEach(fn => {
                fn(eventUpdate)
            });
        }
    }

    function onClose(fn) {
        onUpdate("_close", fn)
    }

    function onOpen(fn) {
        onUpdate("_open", fn)
    }

    function onUpdate(updateType, fn) {
        if (!listener.hasOwnProperty(updateType)) {
            listener[updateType] = []
        }
        listener[updateType].push(fn)
    }
    
    function sendUpdate(eventType, obj) {
        if (!conn) {
            return false;
        }

        conn.send(JSON.stringify({
            name: eventType,
            command: JSON.stringify(obj),
        }))
    }

    return {
        connect,
        onUpdate,
        onClose,
        onOpen,
        sendUpdate,
    }
}