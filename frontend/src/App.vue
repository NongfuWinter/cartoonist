<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';

const showLogin = ref(false)
let timer: any = null
const account = ref('')
const password = ref('')

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
  timer = setInterval(() => {
    console.log(password.value)
  }, 1000)
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
  clearInterval(timer)
})

const isScrolled = ref(false)

const handleScroll = () => {
  isScrolled.value = window.scrollY > 50 // 页面滚动超过 50px 就改变背景
}

function onLeave(_: any, done: any) {
  setTimeout(() => {
    done()
  }, 300)
}
</script>

<template>
  <div :class="[`fixed top-0 left-0 w-full z-50 flex justify-between px-12 h-16 bg-white border-b-1 border-teal-100/50  
    transition-all duration-300 [&>*]:flex`,
    isScrolled ? 'shadow-lg' : 'shadow-sm'
  ]">
    <ul class="flex-none justify-start [&>li]:flex [&>li]:justify-center [&>li]:items-center [&>li]:mr-12">
      <li>logo</li>
      <li>首页</li>
      <li>画廊</li>
      <li>找画师</li>
      <li>黑名单</li>
    </ul>
    <div class="flex-1 basis-128 justify-center items-center">
      <div class="flex-1 flex justify-center items-center max-w-[24rem] min-w-[16rem] h-10 py-3 border-1 border-gray-400 rounded-lg 
        focus-within:max-w-[32rem] focus-within:h-[2.6rem] focus-within:border-teal-500/50 focus-within:ring-1 focus-within:ring-teal-700/20
        transition-all duration-300 ease-in-out">
        <input type="text" class="flex-1 px-3 py-2 border-none" />
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
          class="flex-none size-6 text-gray-400 mr-3">
          <path stroke-linecap="round" stroke-linejoin="round"
            d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
        </svg>
      </div>
    </div>
    <div class="flex-none flex justify-end [&>*]:flex [&>*]:items-center">
      <div class="px-6" @click="showLogin = true">登陆</div>
      <div class="px-6">注册</div>
      <div class="px-4">
        <div class="flex justify-center items-center px-6 py-2 bg-cyan-400/30 border-2 border-cyan-50/80 rounded-2xl shadow-md ring-3 ring-teal-400/20
          transition-transform duration-100 active:scale-95 active:shadow-lg">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="size-4 mr-2 text-teal-700/90">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M3 16.5v2.25A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75V16.5m-13.5-9L12 3m0 0 4.5 4.5M12 3v13.5" />
          </svg>
          <span class="select-none">发布</span>
        </div>

      </div>
    </div>
  </div>
  <div class="flex flex-col h-[100vh] px-36 mt-26 gap-8">
    <div class="p-8 bg-white border-1 border-gray-300 rounded-lg shadow-sm">
      <h1 class="text-xl pl-2 pb-4 border-b-1 border-gray-400">公告</h1>
      <p class="p-2">5.14 更新登陆页面</p>
    </div>
    <div class="p-8 bg-white border-1 border-gray-300 rounded-lg shadow-sm">接单</div>
    <div class="p-8 bg-white border-1 border-gray-300 rounded-lg shadow-sm">作品</div>
    <div class="p-8 bg-white border-1 border-gray-300 rounded-lg shadow-sm">画家</div>
    <div class="p-8 bg-white border-1 border-gray-300 rounded-lg shadow-sm">公告</div>
  </div>
  <transition name="fade" @leave="onLeave">
    <div v-if="showLogin" class="fixed inset-0 z-500 flex items-start justify-center bg-black/50">
      <div class="absolute inset-0"></div>

      <!-- 登录表单 -->
      <div class="login-card relative z-10 flex flex-col bg-white rounded-2xl shadow-lg w-132 mt-32">

        <div class="flex justify-end mx-8 my-10" @click="showLogin = false">
          <div class="p-1 text-gray-500 transition-all duration-200 hover:scale-110 hover:text-red-400">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
              stroke="currentColor" class="size-7">
              <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" />
            </svg>
          </div>
        </div>

        <div class="flex justify-center mx-20 pb-6 border-b-1 border-gray-400">
          <h2 class="text-3xl text-gray-600">登录</h2>
        </div>

        <div class="item mx-2 mt-12 mx-20" style="--i:1">
          <div class="ml-1 mb-2 text-[0.95rem] text-gray-500 transition-all duration-300 ease-out"
            :class="account ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-2'">
            用户名/邮箱
          </div>
          <div class="border-1 border-gray-400 rounded-lg overflow-hidden hover:bg-gray-50">
            <input type="text" v-model="account" placeholder="用户名/邮箱" class="w-full px-4 py-3 text-gray-600" />
          </div>
        </div>

        <div class="item mx-2 mt-6 mx-20" style="--i:2">
          <div class="ml-1 mb-2 text-[0.95rem] text-gray-500 transition-all duration-300 ease-out"
            :class="password ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-2'">
            密码
          </div>
          <div class="border-1 border-gray-400 rounded-lg overflow-hidden hover:bg-gray-50">
            <input type="password" v-model="password" placeholder="密码" class="w-full px-4 py-3 text-gray-600" />
          </div>
        </div>

        <div class="mx-2 mt-6 mx-20 text-[0.93rem] text-right text-sky-600/80">
          <a href="">忘记密码？</a>
        </div>
        <div class="flex justify-center items-center mt-10 mx-20 gap-8">
          <button class="w-full py-[0.8rem] text-[1.1rem] text-gray-50 bg-gradient-to-r from-sky-400 via-blue-400 to-sky-400
             rounded-xl border-1 border-gray-300 shadow-lg transition-all duration-200 hover:brightness-95">
            登录
          </button>
        </div>
        <div class="flex justify-center mt-12 mb-15 text-[0.93rem]">
          <span class="mr-1 text-gray-700">还没有账号? </span>
          <a href="" class="text-sky-600/80">立即注册</a>
        </div>
      </div>
    </div>
  </transition>
</template>

<style scoped lang="css">
:deep(.el-input),
:deep(.el-input__inner),
:deep(.el-textarea__inner) {
  font-size: 16px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity .5s;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* 登录卡片进入 */
.fade-enter-active .login-card {
  animation: popIn 0.5s ease-out forwards;
}

/* 登录卡片退出（倒放） */
.fade-leave-active .login-card {
  animation: popIn 0.5s ease-out reverse forwards;
}

/* 动画 */
@keyframes popIn {
  from {
    transform: translateY(30px);
    opacity: 0;
  }

  to {
    transform: translateY(0);
    opacity: 1;
  }
}

/* ===== stagger ===== */
.item {
  animation: itemIn .5s ease forwards;
  opacity: 0;
  transform: translateY(40px);
  animation-delay: calc(var(--i) * 100ms);
}

@keyframes itemIn {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>