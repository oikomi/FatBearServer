// create Agora client
var client = AgoraRTC.createClient({
  mode: "live",
  codec: "vp8"
});
var localTracks = {
  videoTrack: null,
  audioTrack: null
};
var remoteUsers = {};
// Agora client options
var options = {
  appid: "ac1f26c994ea4f978a11ce1251424920",
  channel: "fatbear",
  uid: "",
  token: "007eJxTYFjDurL39usEP2GDlw8yhc5/MOq428Hw8beP8Lw7JTfzJecrMCQmG6YZmSVbWpqkJpqkWZpbJBoaJqcaGpkamhiZWBoZ7OhWTm0IZGRIKr7FxMgAgSA+O0NaYklSamIRAwMAGLkhlw==",
  role: "audience",
  // host or audience
  audienceLatency: 2
};

// the demo can auto join channel with params in url
// $(() => {
//   var urlParams = new URL(location.href).searchParams;
//   options.appid = urlParams.get("appid");
//   options.channel = urlParams.get("channel");
//   options.token = urlParams.get("token");
//   options.uid = urlParams.get("uid");
//   if (options.appid && options.channel) {
//     $("#uid").val(options.uid);
//     $("#appid").val(options.appid);
//     $("#token").val(options.token);
//     $("#channel").val(options.channel);
//     $("#join-form").submit();
//   }
// });
$("#host-join").click(function (e) {
  options.role = "host";
});
$("#lowLatency").click(function (e) {
  options.role = "audience";
  options.audienceLatency = 1;
  $("#join-form").submit();
});
$("#ultraLowLatency").click(function (e) {
  options.role = "audience";
  options.audienceLatency = 2;
  $("#join-form").submit();
});
$("#join-form").submit(async function (e) {
  e.preventDefault();
  $("#host-join").attr("disabled", true);
  $("#audience-join").attr("disabled", true);
  try {
    // options.channel = $("#channel").val();
    options.uid = Number($("#uid").val());
    // options.appid = $("#appid").val();
    // options.token = $("#token").val();
    await join();
    if (options.role === "host") {
      $("#success-alert a").attr("href", `index.html?appid=${options.appid}&channel=${options.channel}&token=${options.token}`);
      if (options.token) {
        $("#success-alert-with-token").css("display", "block");
      } else {
        $("#success-alert a").attr("href", `index.html?appid=${options.appid}&channel=${options.channel}&token=${options.token}`);
        $("#success-alert").css("display", "block");
      }
    }
  } catch (error) {
    console.error(error);
  } finally {
    $("#leave").attr("disabled", false);
  }
});
$("#leave").click(function (e) {
  leave();
});


$("#send_cmd").click(function (e) {
  send_cmd();
});

async function join() {
  // create Agora client

  if (options.role === "audience") {
    client.setClientRole(options.role, {
      level: options.audienceLatency
    });
    // add event listener to play remote tracks when remote user publishs.
    client.on("user-published", handleUserPublished);
    client.on("user-unpublished", handleUserUnpublished);
  } else {
    client.setClientRole(options.role);
  }

  // join the channel
  options.uid = await client.join(options.appid, options.channel, options.token || null, options.uid || null);
  if (options.role === "host") {
    // create local audio and video tracks
    if (!localTracks.audioTrack) {
      localTracks.audioTrack = await AgoraRTC.createMicrophoneAudioTrack({
        encoderConfig: "music_standard"
      });
    }
    if (!localTracks.videoTrack) {
      localTracks.videoTrack = await AgoraRTC.createCameraVideoTrack();
    }
    // play local video track
    localTracks.videoTrack.play("local-player");
    $("#local-player-name").text(`localTrack(${options.uid})`);
    $("#joined-setup").css("display", "flex");
    // publish local tracks to channel
    await client.publish(Object.values(localTracks));
    console.log("publish success");
  }
}
async function leave() {
  for (trackName in localTracks) {
    var track = localTracks[trackName];
    if (track) {
      track.stop();
      track.close();
      localTracks[trackName] = undefined;
    }
  }

  // remove remote users and player views
  remoteUsers = {};
  $("#remote-playerlist").html("");

  // leave the channel
  await client.leave();
  $("#local-player-name").text("");
  $("#host-join").attr("disabled", false);
  $("#audience-join").attr("disabled", false);
  $("#leave").attr("disabled", true);
  $("#joined-setup").css("display", "none");
  console.log("client leaves channel success");
}

async function send_cmd() {

  axios.post('https://120.55.60.98/cmd', {
    dev_name: '123456',
    cmd: 'do',
    send_user: options.uid
  })
  .then(function (response) {
    alert('send cmd!', 'success')
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });


  // axios.post(`https://120.55.60.98/cmd`,{
  //   params: {
  //     "dev_name": options.devId
  //   }
  // }).then(res=>{
  //   // console.log('res=>',res);

  //   const info = res.data.data.createdAt + " , recive cmd : " +  
  //   res.data.data.cmd + " , to dev "+ res.data.data.dev_name + " ,from " + res.data.data.send_user

  //   const labelcon = document.createTextNode(info);

  //   document.getElementById("area").append(labelcon)
  //   document.getElementById("area").append(document.createTextNode("\n----------------------------\n"))

  //   // document.getElementById("cmd_text").value = JSON.stringify(res.data)         
  // }).catch(err =>{
  //     console.log(err);  //异常的，返回一个错误信息
  // })

  console.log("send_cmd success");
}


async function subscribe(user, mediaType) {
  const uid = user.uid;
  // subscribe to a remote user
  await client.subscribe(user, mediaType);
  console.log("subscribe success");
  if (mediaType === 'video') {
    const player = $(`
      <div id="player-wrapper-${uid}">
        <p class="player-name">remoteUser(${uid})</p>
        <div id="player-${uid}" class="player"></div>
      </div>
    `);
    $("#remote-playerlist").append(player);
    user.videoTrack.play(`player-${uid}`, {
      fit: "contain"
    });
  }
  if (mediaType === 'audio') {
    user.audioTrack.play();
  }
}
function handleUserPublished(user, mediaType) {
  //print in the console log for debugging 
  console.log('"user-published" event for remote users is triggered.');
  const id = user.uid;
  remoteUsers[id] = user;
  subscribe(user, mediaType);
}
function handleUserUnpublished(user, mediaType) {
  //print in the console log for debugging 
  console.log('"user-unpublished" event for remote users is triggered.');
  if (mediaType === 'video') {
    const id = user.uid;
    delete remoteUsers[id];
    $(`#player-wrapper-${id}`).remove();
  }
}