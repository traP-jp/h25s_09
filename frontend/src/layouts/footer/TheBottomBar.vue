<script lang="ts" setup></script>

<template>
    <footer class="$style.bottom-bar">
    <TabBarItem to="/" icon-name="icon-home" label="ホーム" />
    <TabBarItem to="/achievement" icon-name="icon-achievement" label="アチーブメント"/>
    <TabBarItem to="/profile" icon-name="icon-user" label="プロフィール" />
  </footer>
</template>

<style lang="scss" module>
.bottom-bar{
   position: fixed;
   top: auto;
   bottom: 0;
   left: 0;
   width:100%;
   height: 60px;
   padding-bottom: env(safe-area-inset-bottom);
   display: flex;
   justify-content: space-around; 
   align-items: center;
   z-index: 1000;
   background-color: #1f3d1f ;
   border-top: 1px solid #12ae12;
}
</style>