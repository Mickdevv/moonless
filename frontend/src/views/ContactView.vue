<script setup lang="ts">
import { onMounted, ref } from 'vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'
import { useToast } from 'primevue/usetoast'
import 'primeicons/primeicons.css'
import axios from 'axios'
import type { ContactForm } from '@/types/types/contact-form'

const toast = useToast()
const loading = ref(false)

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

const contactFormSubmit = async () => {
  loading.value = true
  try {
    await axios.post('/api/email', {
      name: contactForm.value.name,
      email: contactForm.value.email,
      phone_number: contactForm.value.phoneNumber,
      reason: contactForm.value.reason,
      message: contactForm.value.message,
    })
    contactForm.value = { name: "", reason: "", phoneNumber: "", message: "", email: "" }
    toast.add({ severity: 'success', summary: 'Message sent', detail: "We'll be in touch soon!", life: 4000 })
  } catch {
    toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to send message, please try again.', life: 4000 })
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div>
    <form class="contact-form contact-option" @submit.prevent="contactFormSubmit()">
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
      <Button class="field" type="submit" label="Send message" :loading="loading" />
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
