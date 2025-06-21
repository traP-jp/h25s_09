<script lang="ts" setup>
defineProps<{
  "/": string 
  "/achievement": string
  "/profile": string
  "iconHome" : string
  "iconAchievement": string
  "iconUser": string
  "ホーム": string  
  "アチーブメント": string
  "プロフィール": string
}>()
</script>

<template>
    <router-link :to="to" class="tab-item" active-class="is-active" >
        <i :class="['',iconHome]"></i>
        <span class="">{{"ホーム"}}</span>
    </router-link>
    <router-link :to="achievement" class="tab-item" active-class="is-active" >
        <i :class="['',iconAchievement]"></i>
        <span class="">{{"アチーブメント"}}</span>
    </router-link>
    <router-link :to="profile" class="tab-item" active-class="is-active" >
        <i :class="['',iconUser]"></i>
        <span class="">{{"プロフィール"}}</span>
    </router-link>
</template>

<style lang="scss" module>

.tab-item{
   display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;   
    flex-grow: 1; 
    color: #ff0000; 
    text-decoration: none;
}
.icon {
  font-size: 240px;
}
.label {
  font-size: 10px;
}
.is-active{
    color: #ff0000;
    font-weight: 600;
}
</style>
