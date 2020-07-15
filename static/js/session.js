

// console.log("session script")

// console.log(document.cookie)

if(document.cookie != "") {
    let logoutButton = document.getElementById("logoutButton")

    logoutButton.addEventListener("click", (event) => {
        console.log("logout button pressed")
        document.cookie = "user_uuid=; expires=Thu, 01 Jan 1970 00:00:01 GMT; max-age=-1;"
        document.cookie = "username=; expires=Thu, 01 Jan 1970 00:00:01 GMT; max-age=-1;"
        location.reload()
        //document.cookie = "user_uuid="
    })
}


