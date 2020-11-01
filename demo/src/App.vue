<style lang="sass" scoped>
@font-face
  font-family: roboto
  src: url(/fonts.ttf)
@font-face
  font-family: roboto-bold
  src: url(/fonts-bold.ttf)
.content
  width: 100%
  height: 100%
  box-sizing: border-box
  display: flex
  flex-direction: column
  align-items: center
  justify-content: center
.scrollable
  overflow-x: hidden
  overflow-y: auto
.title
  font-size: 2em
  font-weight: bold
  margin-top: 32px
  font-family: roboto-bold
.subtitle
  font-size: 1.5em
  margin-bottom: 32px
  font-family: roboto
.picture
  width: 75vw
  height: 50vh
  background-color: lightgray
  border-radius: 32px
  object-fit: cover
.description
  width: 75vw
  text-align: center
  margin: 32px
  font-family: roboto
.footer
  width: 100%
  box-sizing: border-box
  padding: 8px
  text-align: center
  color: white
  background-color: black
  font-family: roboto
</style>

<template>
  <div class="content">
    <div class="content scrollable">
      <div class="title">Kuda Web Server</div>
      <div class="subtitle">Demonstration</div>
      <img
        class="picture"
        :src="`/img/${img_id}.jpg?t=${randomNumber()}`"
        alt="Landscape pictures at Lombok Island"
        @load="nextPicture"
      />
      <p class="description">
        Kuda web server specialized for serving SPA front-end.
        It gives flexibility to developer for taking care router
        and rendering on client side. So all assets can be stored
        in server RAM. Thus, a lot faster loading time.
      </p>
    </div>
    <div class="footer">
      © Copyright 2020 Athaariq Ardhiansyah, All Pictures on this Site are Right Reserved.
    </div>
  </div>
</template>

<script>
let lastTimeout = NaN;

export default {
  data() {
    return {
      img_id: 1
    }
  },
  methods: {
    randomNumber() {
      return Math.round(Math.random() * Number.MAX_SAFE_INTEGER);
    },
    nextPicture() {
      if(!isNaN(lastTimeout)) {
        clearTimeout(lastTimeout);
        lastTimeout = NaN;
      }
      lastTimeout = setTimeout(() => {
        this.img_id++;
        if(this.img_id > 15) {
          this.img_id -= 15
        }
      },2000);
    }
  }
}
</script>
