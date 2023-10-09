

import AgoraRTC from "agora-rtc-sdk-ng";
import axios from 'axios'


let rtc = {
    localAudioTrack: null,
    localVideoTrack: null,
    client: null,
};

let options = {
    // Pass your App ID here.
    appId: "ac1f26c994ea4f978a11ce1251424920",
    // Set the channel name.
    channel: "fatbear",
    // Pass your temp token here.
    token: "007eJxTYFjC/o6v16OjYv/e1ttRxXd2935Iinj6+Mr+JFXDRu8S8y8KDInJhmlGZsmWliapiSZpluYWiYaGyamGRqaGJkYmlkYGF+ZLpzYEMjIEdzMxMjJAIIjPzpCWWJKUmljEwAAAMVEhYw==",
    // Set the user ID.
    uid: 123456,

    devId: "0",
};


// var int=self.setInterval("clock()",1000);
function get(url,params={}) {
    return new Promise(
        //一个正常的结果，一个异常的结果
        (resolve,reject) =>{
            axios.get(url,{
                headers: {
                  'Access-Control-Allow-Origin': '*'
                },
                crossdomain: true
              })
                .then(response=>{
                    console.log(response.data)

                    resolve(response.data);   //正常的，返回响应数据
                })
                .catch(
                    err =>{
                        reject(err);  //异常的，返回一个错误信息
                    }
                )
        }
    );

}


function fun() {
    console.log("start timer")
    // const config = {
    //     headers:{
    //       "Access-Control-Allow-Origin": "*",
    //       "Access-Control-Allow-Headers": "X-Requested-With,Content-Type",
    //       "Access-Control-Allow-Methods": "PUT,POST,GET,DELETE,OPTIONS"
    //     }
    //   };
      
    // let data = {"dev_name":"123456","cmd":"get_cmd"};
    // axios.get(`http://127.0.0.1:8080/api/v1/dev/cmd`, config).then(res=>{
    //     console.log('res=>',res);            
    // })
    // get("http://127.0.0.1:8080/api/v1/dev/cmd", {})

    if (options.devId == "0" || options.devId == "") {
        return
    }

    axios.get(`http://127.0.0.1:8080/api/v1/dev/cmd`,{
        params: {
          "dev_name": options.devId
        }
      }).then(res=>{
        // console.log('res=>',res);

        const info = res.data.data.createdAt + " , recive cmd : " +  
        res.data.data.cmd + " , to dev "+ res.data.data.dev_name + " ,from " + res.data.data.send_user

        const labelcon = document.createTextNode(info);

        document.getElementById("area").append(labelcon)
        document.getElementById("area").append(document.createTextNode("\n----------------------------\n"))

        // document.getElementById("cmd_text").value = JSON.stringify(res.data)         
    }).catch(err =>{
        console.log(err);  //异常的，返回一个错误信息
    }
)

    //console.log("time end")
}
setInterval(fun, 10000);

async function startBasicCall() {
    // Create an AgoraRTCClient object.
    rtc.client = AgoraRTC.createClient({mode: "rtc", codec: "vp8"});

    // Listen for the "user-published" event, from which you can get an AgoraRTCRemoteUser object.
    rtc.client.on("user-published", async (user, mediaType) => {
        // Subscribe to the remote user when the SDK triggers the "user-published" event
        await rtc.client.subscribe(user, mediaType);
        console.log("subscribe success");

        // If the remote user publishes a video track.
        if (mediaType === "video") {
            // Get the RemoteVideoTrack object in the AgoraRTCRemoteUser object.
            const remoteVideoTrack = user.videoTrack;
            // Dynamically create a container in the form of a DIV element for playing the remote video track.
            const remotePlayerContainer = document.createElement("div");
            // Specify the ID of the DIV container. You can use the uid of the remote user.
            remotePlayerContainer.id = user.uid.toString();
            remotePlayerContainer.textContent = "Remote user " + user.uid.toString();
            remotePlayerContainer.style.width = "640px";
            remotePlayerContainer.style.height = "480px";
            document.body.append(remotePlayerContainer);

            // Play the remote video track.
            // Pass the DIV container and the SDK dynamically creates a player in the container for playing the remote video track.
            remoteVideoTrack.play(remotePlayerContainer);
        }

        // If the remote user publishes an audio track.
        if (mediaType === "audio") {
            // Get the RemoteAudioTrack object in the AgoraRTCRemoteUser object.
            const remoteAudioTrack = user.audioTrack;
            // Play the remote audio track. No need to pass any DOM element.
            remoteAudioTrack.play();
        }

        // Listen for the "user-unpublished" event
        rtc.client.on("user-unpublished", user => {
            // Get the dynamically created DIV container.
            const remotePlayerContainer = document.getElementById(user.uid);
            // Destroy the container.
            remotePlayerContainer.remove();
        });
    });

    window.onload = function () {
        // document.getElementById("join").onclick = async function () {
        //     // Join an RTC channel.
        //     await rtc.client.join(options.appId, options.channel, options.token, options.uid);
        //     // Create a local audio track from the audio sampled by a microphone.
        //     rtc.localAudioTrack = await AgoraRTC.createMicrophoneAudioTrack();
        //     // Create a local video track from the video captured by a camera.
        //     rtc.localVideoTrack = await AgoraRTC.createCameraVideoTrack();
        //     // Publish the local audio and video tracks to the RTC channel.
        //     await rtc.client.publish([rtc.localAudioTrack, rtc.localVideoTrack]);
        //     // Dynamically create a container in the form of a DIV element for playing the local video track.
        //     const localPlayerContainer = document.createElement("div");
        //     // Specify the ID of the DIV container. You can use the uid of the local user.
        //     localPlayerContainer.id = options.uid;
        //     localPlayerContainer.textContent = "Local user " + options.uid;
        //     localPlayerContainer.style.width = "640px";
        //     localPlayerContainer.style.height = "480px";
        //     document.body.append(localPlayerContainer);

        //     // Play the local video track.
        //     // Pass the DIV container and the SDK dynamically creates a player in the container for playing the local video track.
        //     rtc.localVideoTrack.play(localPlayerContainer);
        //     console.log("publish success!");
        // };

        document.getElementById("join").onclick = async function () {
            // Join an RTC channel.
            await rtc.client.join(options.appId, options.channel, options.token, options.uid);
            // Create a local audio track from the audio sampled by a microphone.
            rtc.localAudioTrack = await AgoraRTC.createMicrophoneAudioTrack();
            // Create a local video track from the video captured by a camera.
            rtc.localVideoTrack = await AgoraRTC.createCameraVideoTrack();
            // Publish the local audio and video tracks to the RTC channel.
            await rtc.client.publish([rtc.localAudioTrack, rtc.localVideoTrack]);
            // Dynamically create a container in the form of a DIV element for playing the local video track.
            // const localPlayerContainer = document.createElement("div");
            // // Specify the ID of the DIV container. You can use the uid of the local user.
            // localPlayerContainer.id = options.uid;
            // localPlayerContainer.textContent = "Local user " + options.uid;
            // localPlayerContainer.style.width = "640px";
            // localPlayerContainer.style.height = "480px";
            // document.body.append(localPlayerContainer);

            const localPlayerContainer = document.getElementById("video_info")

            localPlayerContainer.style.width = "640px";
            localPlayerContainer.style.height = "480px";

            // Play the local video track.
            // Pass the DIV container and the SDK dynamically creates a player in the container for playing the local video track.
            rtc.localVideoTrack.play(localPlayerContainer);
            console.log("publish success!");


        };


        document.getElementById("leave").onclick = async function () {
            // Destroy the local audio and video tracks.
            rtc.localAudioTrack.close();
            rtc.localVideoTrack.close();

            // Traverse all remote users.
            rtc.client.remoteUsers.forEach(user => {
                // Destroy the dynamically created DIV containers.
                const playerContainer = document.getElementById(user.uid);
                playerContainer && playerContainer.remove();
            });

            // Leave the channel.
            await rtc.client.leave();
        };


        document.getElementById("login").onclick = async function () {

            const devId = document.getElementById("device_id_input").value

            options.devId = devId
            console.log(devId)
            console.log("options.devId: " , options.devId)

        };


        // document.getElementById("do").onclick = async function () {

        //     // let data = {"cmd":"xxx","name":"yyyy"};
        //     // axios.post(`${this.$url}/test/testRequest`,data).then(res=>{
        //     //     console.log('res=>',res);            
        //     // })


        //     const cmdContainer = document.createElement("div");
        //     // Specify the ID of the DIV container. You can use the uid of the local user.
        //     cmdContainer.id = options.uid;
        //     cmdContainer.textContent = "receive command " + options.uid;
        //     cmdContainer.style.width = "640px";
        //     cmdContainer.style.height = "480px";
        //     document.body.append(cmdContainer);

        //     console.log("do success!");
        // };


    };
}

startBasicCall();
