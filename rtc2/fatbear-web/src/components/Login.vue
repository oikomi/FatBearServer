<script setup lang="ts">

import { inject, ref } from 'vue'
import { useRouter } from 'vue-router'

import { useTokenStore } from '@/stores/token';

import { userStore } from '@/stores/user';
import {APP_ID, TOKEN, SERVER_BASE} from '@/config/config';

const tokenStore = useTokenStore()

const uStore = userStore()

// dev
const USER_LOOGIN_URL = SERVER_BASE + "api/v1/user/login"

const axios: any = inject('axios')  // inject axios

const router = useRouter()

const userName = ref('')
const password = ref('')
const channel = ref('')
let host = ""
const radioUser = ref('')

async function doLogin() {
	console.log("start login")

	console.log(userName)
	console.log(password)

	console.log("radioUser : ", radioUser.value)

	if (radioUser.value === 'user') {
		host = 'user'
	} else if (radioUser.value === 'model') {
		host = 'host'
	}

	console.log("host is ", host)

	axios
		.post(USER_LOOGIN_URL, {
			'user_name': userName.value,
			'password': password.value,
			'role': host
		})
		.then((response: { data: any }) => {
			console.log(response.data)
			if (response.data.code !== 0) {
				console.log("toast come")
				alert("login failed")
			} else {
				console.log("login success ,push to live")
				tokenStore.setToken(response.data.data)
				uStore.setHost(host)
				uStore.setUserName(userName.value)
				console.log("token is  :", tokenStore.getToken())
				console.log("host is  :", uStore.getHost())

				if (uStore.getHost() === 'host') {
					router.push({ name: 'live' })
				} else {
					router.push({ name: 'model' })
				}
			}
		}).catch((err: any) => {
			console.log(err);
			alert("[ex]login failed")
		});
}


</script>

<template>
	<body class="d-flex align-items-center py-4">
		<svg xmlns="http://www.w3.org/2000/svg" class="d-none">
			<symbol id="check2" viewBox="0 0 16 16">
				<path
					d="M13.854 3.646a.5.5 0 0 1 0 .708l-7 7a.5.5 0 0 1-.708 0l-3.5-3.5a.5.5 0 1 1 .708-.708L6.5 10.293l6.646-6.647a.5.5 0 0 1 .708 0z" />
			</symbol>
			<symbol id="circle-half" viewBox="0 0 16 16">
				<path d="M8 15A7 7 0 1 0 8 1v14zm0 1A8 8 0 1 1 8 0a8 8 0 0 1 0 16z" />
			</symbol>
			<symbol id="moon-stars-fill" viewBox="0 0 16 16">
				<path
					d="M6 .278a.768.768 0 0 1 .08.858 7.208 7.208 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277.527 0 1.04-.055 1.533-.16a.787.787 0 0 1 .81.316.733.733 0 0 1-.031.893A8.349 8.349 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.752.752 0 0 1 6 .278z" />
				<path
					d="M10.794 3.148a.217.217 0 0 1 .412 0l.387 1.162c.173.518.579.924 1.097 1.097l1.162.387a.217.217 0 0 1 0 .412l-1.162.387a1.734 1.734 0 0 0-1.097 1.097l-.387 1.162a.217.217 0 0 1-.412 0l-.387-1.162A1.734 1.734 0 0 0 9.31 6.593l-1.162-.387a.217.217 0 0 1 0-.412l1.162-.387a1.734 1.734 0 0 0 1.097-1.097l.387-1.162zM13.863.099a.145.145 0 0 1 .274 0l.258.774c.115.346.386.617.732.732l.774.258a.145.145 0 0 1 0 .274l-.774.258a1.156 1.156 0 0 0-.732.732l-.258.774a.145.145 0 0 1-.274 0l-.258-.774a1.156 1.156 0 0 0-.732-.732l-.774-.258a.145.145 0 0 1 0-.274l.774-.258c.346-.115.617-.386.732-.732L13.863.1z" />
			</symbol>
			<symbol id="sun-fill" viewBox="0 0 16 16">
				<path
					d="M8 12a4 4 0 1 0 0-8 4 4 0 0 0 0 8zM8 0a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 0zm0 13a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-1 0v-2A.5.5 0 0 1 8 13zm8-5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 .5.5zM3 8a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1 0-1h2A.5.5 0 0 1 3 8zm10.657-5.657a.5.5 0 0 1 0 .707l-1.414 1.415a.5.5 0 1 1-.707-.708l1.414-1.414a.5.5 0 0 1 .707 0zm-9.193 9.193a.5.5 0 0 1 0 .707L3.05 13.657a.5.5 0 0 1-.707-.707l1.414-1.414a.5.5 0 0 1 .707 0zm9.193 2.121a.5.5 0 0 1-.707 0l-1.414-1.414a.5.5 0 0 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .707zM4.464 4.465a.5.5 0 0 1-.707 0L2.343 3.05a.5.5 0 1 1 .707-.707l1.414 1.414a.5.5 0 0 1 0 .708z" />
			</symbol>
		</svg>


		<main class="form-signin w-100 m-auto maincss rounded-3">
			<img class="mb-3" src="/src/assets/SVG/SVG/huatong.svg" alt="" width="72" height="57">
			<h1 class="h3 mb-3 fw-normal fw-bold">Livestream Demo System</h1>

			<div class="row mb-3">
			<img class="col-2 " src="/src/assets/SVG/SVG/shouzhang.svg" alt="" width="46" height="35">
			<h4 class="col-10 h5  fw-normal">Just for demo Using</h4>
			</div>

			<div class="form-floating">
				<input type="text" class="form-control" id="user_id" placeholder="123456" v-model="userName">
				<label for="user_id">User ID</label>
			</div>
			<div class="form-floating">
				<input type="password" class="form-control" id="password" placeholder="Password" v-model="password">
				<label for="password">Password</label>
			</div>

			<!-- <div class="form-floating mt-0">
				<input type="text" class="form-control" id="channel" placeholder="channel" v-model="channel">
				<label for="channel">Channel</label>
			</div> -->

			<div class="form-check mt-2">
				<div class="row">
					<div class="col form-check">
						<input class="form-check-input" type="radio" name="flexRadioDefault" v-model="radioUser"
							id="flexRadioDefault1"  value="user">
						<label class="form-check-label" for="flexRadioDefault1">
							User
						</label>
					</div>
					<div class="col form-check">
						<input class="form-check-input" type="radio" name="flexRadioDefault" v-model="radioUser"
							id="flexRadioDefault2" value="model">
						<label class="form-check-label" for="flexRadioDefault2">
							Model
						</label>
					</div>
				</div>

			</div>

			<!-- <div class="form-check text-start my-3">
				<input class="form-check-input" type="checkbox" value="remember-me" id="flexCheckDefault">
				<label class="form-check-label" for="flexCheckDefault">
					Remember me
				</label>
			</div> -->
			<button class="btn btn-primary w-100 py-2 mt-2 secai" @click="doLogin">Login</button>
			<p class="mt-5 mb-3 text-body-secondary">&copy; 2023</p>
		</main>

		<!-- <button class="btn btn-primary w-100 py-2" @click="doLogin">Sign in</button> -->


		

	</body>
</template>

<style scoped>
html,
body {
	width : 100%;
    height : 100%;
	/* background-color:#C24362; */
	background-image: linear-gradient(-45deg, #C24362, #6450A4);
}

.secai {
	background-image: linear-gradient(-45deg, #C24362, #6450A4);
}

.maincss {
	background-color:#FBFCF9;
}

.form-signin {
	max-width: 330px;
	padding: 1rem;
}

.form-signin .form-floating:focus-within {
	z-index: 2;
}

.form-signin input[type="email"] {
	margin-bottom: -1px;
	border-bottom-right-radius: 0;
	border-bottom-left-radius: 0;
}

.form-signin input[type="password"] {
	margin-bottom: 10px;
	border-top-left-radius: 0;
	border-top-right-radius: 0;
}

.bd-placeholder-img {
	font-size: 1.125rem;
	text-anchor: middle;
	-webkit-user-select: none;
	-moz-user-select: none;
	user-select: none;
}

@media (min-width: 768px) {
	.bd-placeholder-img-lg {
		font-size: 3.5rem;
	}
}

.b-example-divider {
	width: 100%;
	height: 3rem;
	background-color: rgba(0, 0, 0, .1);
	border: solid rgba(0, 0, 0, .15);
	border-width: 1px 0;
	box-shadow: inset 0 .5em 1.5em rgba(0, 0, 0, .1), inset 0 .125em .5em rgba(0, 0, 0, .15);
}

.b-example-vr {
	flex-shrink: 0;
	width: 1.5rem;
	height: 100vh;
}

.bi {
	vertical-align: -.125em;
	fill: currentColor;
}

.nav-scroller {
	position: relative;
	z-index: 2;
	height: 2.75rem;
	overflow-y: hidden;
}

.nav-scroller .nav {
	display: flex;
	flex-wrap: nowrap;
	padding-bottom: 1rem;
	margin-top: -1px;
	overflow-x: auto;
	text-align: center;
	white-space: nowrap;
	-webkit-overflow-scrolling: touch;
}

.btn-bd-primary {
	--bd-violet-bg: #712cf9;
	--bd-violet-rgb: 112.520718, 44.062154, 249.437846;

	--bs-btn-font-weight: 600;
	--bs-btn-color: var(--bs-white);
	--bs-btn-bg: var(--bd-violet-bg);
	--bs-btn-border-color: var(--bd-violet-bg);
	--bs-btn-hover-color: var(--bs-white);
	--bs-btn-hover-bg: #6528e0;
	--bs-btn-hover-border-color: #6528e0;
	--bs-btn-focus-shadow-rgb: var(--bd-violet-rgb);
	--bs-btn-active-color: var(--bs-btn-hover-color);
	--bs-btn-active-bg: #5a23c8;
	--bs-btn-active-border-color: #5a23c8;
}

.bd-mode-toggle {
	z-index: 1500;
}

.bd-mode-toggle .dropdown-menu .active .bi {
	display: block !important;
}
</style>
