

let messageForm = document.getElementById("inputForm")
let sendButton = document.getElementById("sendButton")

// let messages = document.getElementById("chatbox")

// let session = document.cookie
// console.log(document.cookie)

let splittedCookie = document.cookie.split(";")

// 10 - len to uuid <user_uuid=...>
let user_uuid = splittedCookie[0].slice(10)

// 10 - len to username <username=...>
let username = splittedCookie[1].slice(10)

// console.log(user_uuid)
// console.log(username)


let ws = new WebSocket("ws://localhost:8080/ws")


ws.addEventListener("open", (event) => {
    console.log("connected")

    sendButton.addEventListener("click", (event) => {
        event.preventDefault()
        
        let messageText = messageForm.value
        messageForm.value = ""
    
        let data = {
            "username": username,
            "user_uuid": user_uuid,
            "text": messageText
        }
        
        ws.send(JSON.stringify(data))
    })
})

// jquery
var $messages = $("#chatbox")

ws.addEventListener("message", (event) => {
    // view in form
    // let message = event.data
    //console.log(event.data)
    let message = JSON.parse(event.data)
    // console.log(message)
    // console.log(message["username"] + " -> " + message["text"])

    console.log(message["username"], "->", message["text"])

    // view message in chatbox
    $messages.append("<div>" + message["username"] + " -> " + message["text"] + "</div>")
})

ws.addEventListener("close", (event) => {
    console.log(event.data)
    console.log("websocket closed")
})

ws.addEventListener("error", (event) => {
    console.log(event.data)
    console.log("error!!!")
})