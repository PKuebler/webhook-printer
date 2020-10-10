window.onload = function () {
    infoBody = document.getElementById("info-body")
    inspectorBody = document.getElementById("inspector-body")
    const ws = NewClient()

    ws.onUpdate('config', (evt) => {
        console.log(evt)

        infoBody.innerHTML = "Use this Webhook URL to test your System:<br />" + document.location.host + "/hook/" + evt.hookID
    })

    ws.onUpdate('incoming', (evt) => {
        console.log(evt)

        const time = document.createElement("div")
        time.classList.add("entry-time")
        time.innerHTML = new Date()

        const method = document.createElement("div")
        method.classList.add("entry-method")
        method.innerHTML = "/"

        const head = document.createElement("div")
        head.classList.add("entry-head")
        head.appendChild(time)
        head.appendChild(method)

        const code = document.createElement("code")
        code.classList.add("json")
        code.innerHTML = JSON.stringify(evt, null, 4)

        const body = document.createElement("pre")
        body.classList.add("entry-body")
        body.appendChild(code)

        const entry = document.createElement("div")
        entry.classList.add("entry")
        entry.appendChild(head)
        entry.appendChild(body)

        inspectorBody.appendChild(entry)

        hljs.highlightBlock(code);
    })

    ws.onOpen(() => {
        console.log("open")
    })

    ws.connect(document.location.host)
}