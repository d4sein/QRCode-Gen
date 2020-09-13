<template>
  <div class="container">
    <form v-on:submit.prevent="submitForm">
      <div class="field">
        <label class="label">Size</label>
        <div class="control">
          <input v-model="data.size"
          class="input is-small is-primary"
          type="text" placeholder="256">
        </div>
        <p class="help">The image will be automatically adjusted if size is too small</p>
      </div>
      <div class="field">
        <label class="label">Content</label>
        <div class="control">
          <textarea v-model="data.content"
          class="textarea is-primary is-small has-fixed-size"
          placeholder="Content"></textarea>
        </div>
      </div>
      <div class="control">
        <button class="button is-primary">Submit</button>
      </div>
    </form>
    <div class="qrcode">
      <img :src="data.base64 + data.png" alt="">
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import ky from 'ky'

interface ResponseStruct {
  content: string;
}

@Component
export default class Container extends Vue {
  data = {
    size: '',
    content: '',
    base64: 'data:image/png;base64, ',
    png: ''
  }
  submitForm = async () => {
    const data = {
      content: this.data.content,
      size: parseInt(this.data.size) || 256
    }

    const res: ResponseStruct = await ky.post('http://localhost:3000/', {json: data}).json()
    this.data.png = res['content']
  }
}
</script>

<style lang="less">
.qrcode {
  display: block;
  margin-left: auto;
  margin-right: auto;
  width: 50%;
}

@media only screen and (max-width: 800px) {
  .container {
    width: 80% !important;
    padding-top: 50px;
  }
}

@media only screen and (min-width: 800px) {
  .container {
    width: 50% !important;
    padding-top: 50px;
  }
}
</style>
