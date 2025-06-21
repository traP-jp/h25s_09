<script lang="ts" setup></script>

<template>
    <footer class="bottom-bar">
    <TabBarItem to="/" iconName="iconHome" label="ホーム" />
    <TabBarItem to="/achievement" iconName="iconAchievement" label="アチーブメント"/>
    <TabBarItem to="/profile" iconName="iconUser" label="プロフィール" />
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
   display: flex;
   justify-content: space-around; 
   align-items: center;
   z-index: 1000;
   background-color: #1f3d1f ;
   border-top: 1px solid #12ae12;
}
</style>