<script setup lang="ts">
import { onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useProductStore } from '@/stores/products'
import InputText from 'primevue/inputtext'
import InputNumber from 'primevue/inputnumber'
import Checkbox from 'primevue/checkbox'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'
import ConfirmPopup from 'primevue/confirmpopup';


import { useConfirm } from "primevue/useconfirm";
import { useToast } from "primevue/usetoast";
import 'primeicons/primeicons.css'
import { ref, nextTick } from 'vue'
import type { CreateProductImageDto } from '@/types/DTOs/CreateProductImage.dto'
import { useAuthStore } from '@/stores/auth'
import type { ContactForm } from '@/types/types/contact-form'


const route = useRoute()
const productStore = useProductStore()
const authStore = useAuthStore()

const confirm = useConfirm();
const toast = useToast();

onMounted(async () => {
  let tiktokScript = document.createElement('script')
  tiktokScript.setAttribute('src', 'https://www.tiktok.com/embed.js')
  document.head.appendChild(tiktokScript)

  let instagramScript = document.createElement('script')
  instagramScript.setAttribute('src', "//www.instagram.com/embed.js")
  document.head.appendChild(instagramScript)

})

const contactForm = ref<ContactForm>({
  name: "",
  reason: "",
  phoneNumber: "",
  message: "",
  email: ""
})

const contactFormSubmit = () => {
  console.log("TODO: Submit contact form")
}


</script>

<template>
  <h2>Contact options</h2>
  <div>
    <form class="contact-form contact-option" @submit="contactFormSubmit()">
      <div class="field">
        <label>Name</label>
        <InputText required v-model="contactForm.name" class="w-full" />
      </div>
      <div class="field">
        <label>Email</label>
        <InputText required v-model="contactForm.email" class="w-full" />
      </div>
      <div class="field">
        <label>Phone number</label>
        <InputText required v-model="contactForm.phoneNumber" class="w-full" />
      </div>
      <div class="field">
        <label>Reason</label>
        <InputText required v-model="contactForm.reason" class="w-full" />
      </div>
      <div class="field">
        <label>Message</label>
        <Textarea required v-model="contactForm.message" class="w-full" />
      </div>
    </form>

    <div class="contact-option">

      <h3>Send us an email at <a href="mailto:bandmoonless@gmail.com">bandmoonless@gmail.com</a>
      </h3>

    </div>

    <div class="contact-option">
      <h3>DM us on social media!</h3>
      <iframe src="https://www.Instagram.com/moonlessoff/embed" width="100%" frameborder="0" scrolling="no"
        allowtransparency="true" style="border: none; overflow: hidden;" class="instagram-embed">
      </iframe>

      <blockquote class="tiktok-embed" cite="https://www.tiktok.com/@moonlessoff" data-unique-id="moonlessoff"
        data-embed-type="creator">
        <section> <a target="_blank" href="https://www.tiktok.com/@moonlessoff?refer=creator_embed">@moonlessoff</a>
        </section>
      </blockquote>
    </div>
  </div>
</template>

<style>
Button {
  margin: 2rem;
}

@media (min-width: 800px) {
  .contact-option {
    width: 700px;
  }
}

.contact-option {
  padding-top: 2rem;
  padding-bottom: 2rem;
  border-top: solid;
}

.tiktok-embed .instagram-embed {
  border-radius: 1rem;
}

.contact-form {
  display: flex;
  flex-direction: column;
}

.field {
  justify-content: space-between;
  display: flex;
  margin: 0.5rem;
}


h2 {
  width: 100%;
  padding-bottom: 2rem;
}
</style>
