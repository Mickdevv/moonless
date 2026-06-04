<script setup lang="ts">
import { nextTick, onMounted, ref } from 'vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Textarea from 'primevue/textarea'
import { useToast } from 'primevue/usetoast'
import 'primeicons/primeicons.css'
import axios from 'axios'
import type { ContactForm } from '@/types/types/contact-form'

const toast = useToast()
const loading = ref(false)
const embedKey = ref(0)

const loadTikTokScript = () => {
  document.querySelector('script[src="https://www.tiktok.com/embed.js"]')?.remove()
  const script = document.createElement('script')
  script.src = 'https://www.tiktok.com/embed.js'
  document.head.appendChild(script)
}

const retryEmbeds = async () => {
  embedKey.value++
  await nextTick()
  loadTikTokScript()
}

onMounted(() => {
  loadTikTokScript()
})

const contactForm = ref<ContactForm>({
  name: '',
  reason: '',
  phoneNumber: '',
  message: '',
  email: '',
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
    contactForm.value = { name: '', reason: '', phoneNumber: '', message: '', email: '' }
    toast.add({ severity: 'success', summary: 'Message sent', detail: "We'll be in touch soon!", life: 4000 })
  } catch {
    toast.add({ severity: 'error', summary: 'Error', detail: 'Failed to send message, please try again.', life: 4000 })
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <section class="contact-section">
    <div class="section-header">
      <h1 class="section-title">CONTACT</h1>
      <div class="title-accent"></div>
    </div>

    <div class="contact-layout">

      <!-- Left: Contact form -->
      <div class="contact-form-col">
        <h2 class="col-heading">Send us a message</h2>
        <form class="contact-form" @submit.prevent="contactFormSubmit">
          <div class="field">
            <label class="field-label">Name</label>
            <InputText required v-model="contactForm.name" class="field-input" placeholder="Your name" />
          </div>
          <div class="field">
            <label class="field-label">Email</label>
            <InputText required v-model="contactForm.email" type="email" class="field-input" placeholder="your@email.com" />
          </div>
          <div class="field">
            <label class="field-label">Phone number</label>
            <InputText v-model="contactForm.phoneNumber" class="field-input" placeholder="+33 6 00 00 00 00" />
          </div>
          <div class="field">
            <label class="field-label">Reason</label>
            <InputText required v-model="contactForm.reason" class="field-input" placeholder="Booking, press, collaboration…" />
          </div>
          <div class="field">
            <label class="field-label">Message</label>
            <Textarea required v-model="contactForm.message" class="field-input field-textarea" :autoResize="false" rows="5" placeholder="Tell us more…" />
          </div>
          <Button
            type="submit"
            label="Send message"
            :loading="loading"
            class="submit-btn"
          />
        </form>
      </div>

      <!-- Right: Other contact options -->
      <div class="contact-info-col">
        <h2 class="col-heading">Other ways to reach us</h2>

        <a href="mailto:bandmoonless@gmail.com" class="email-link">
          <i class="pi pi-envelope"></i>
          <span>bandmoonless@gmail.com</span>
        </a>

        <div class="social-section">
          <p class="social-label">Find us on social media</p>

          <div :key="embedKey" class="embeds">
            <div class="embed-wrapper">
              <iframe
                src="https://www.Instagram.com/moonlessoff/embed"
                width="100%"
                frameborder="0"
                scrolling="no"
                allowtransparency="true"
                class="instagram-embed"
              ></iframe>
            </div>

            <div class="embed-wrapper tiktok-wrapper">
              <blockquote
                class="tiktok-embed"
                cite="https://www.tiktok.com/@moonlessoff"
                data-unique-id="moonlessoff"
                data-embed-type="creator"
              >
                <section>
                  <a target="_blank" href="https://www.tiktok.com/@moonlessoff?refer=creator_embed">@moonlessoff</a>
                </section>
              </blockquote>
            </div>
          </div>

          <button class="reload-btn" @click="retryEmbeds()">
            <i class="pi pi-refresh"></i> Reload embeds
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.contact-section {
  width: 100%;
  padding: 4rem 2rem 6rem;
  background: var(--band-black);
  min-height: 100vh;
}

/* ── Two-column layout ── */
.contact-layout {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  max-width: 1100px;
  margin: 0 auto;
  align-items: start;
}

.col-heading {
  font-size: 0.75rem;
  font-weight: 700;
  letter-spacing: 0.2em;
  text-transform: uppercase;
  color: var(--band-white-muted);
  margin-bottom: 1.75rem;
  padding-bottom: 0.6rem;
  border-bottom: 1px solid var(--color-border);
}

/* ── Form ── */
.contact-form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.field-label {
  font-size: 0.75rem;
  font-weight: 600;
  letter-spacing: 0.1em;
  text-transform: uppercase;
  color: var(--band-white-muted);
}

/* Override PrimeVue InputText */
:deep(.field-input.p-inputtext),
:deep(.field-input.p-textarea) {
  background: var(--band-black-card) !important;
  border: 1px solid var(--color-border) !important;
  border-radius: 0 !important;
  color: var(--band-white) !important;
  padding: 0.65rem 0.85rem !important;
  font-size: 0.9rem !important;
  width: 100%;
  transition: border-color 0.2s !important;
  box-shadow: none !important;
}

:deep(.field-input.p-inputtext:focus),
:deep(.field-input.p-textarea:focus) {
  border-color: var(--band-yellow) !important;
  outline: none !important;
  box-shadow: 0 0 0 2px rgba(245, 197, 24, 0.15) !important;
}

:deep(.field-input.p-inputtext::placeholder),
:deep(.field-input.p-textarea::placeholder) {
  color: var(--band-white-muted) !important;
  opacity: 0.6;
}

.field-textarea {
  resize: vertical;
  min-height: 120px;
}

/* Submit button */
:deep(.submit-btn.p-button) {
  background: var(--band-yellow) !important;
  border: none !important;
  border-radius: 0 !important;
  color: var(--band-black) !important;
  font-size: 0.8rem !important;
  font-weight: 700 !important;
  letter-spacing: 0.18em !important;
  text-transform: uppercase !important;
  padding: 0.85rem 2rem !important;
  margin-top: 0.5rem;
  transition: background 0.2s, transform 0.2s !important;
  box-shadow: none !important;
  width: 100%;
  justify-content: center;
}

:deep(.submit-btn.p-button:hover) {
  background: var(--band-yellow-dark) !important;
  transform: translateY(-2px);
}

:deep(.submit-btn.p-button:focus) {
  box-shadow: 0 0 0 2px rgba(245, 197, 24, 0.3) !important;
}

/* ── Right column ── */
.email-link {
  display: inline-flex;
  align-items: center;
  gap: 0.6rem;
  color: var(--band-yellow);
  font-size: 0.95rem;
  font-weight: 600;
  text-decoration: none;
  padding: 0.85rem 1rem;
  border: 1px solid var(--color-border);
  background: var(--band-black-card);
  margin-bottom: 2rem;
  width: 100%;
  transition: border-color 0.2s, background 0.2s;
}

.email-link:hover {
  border-color: var(--band-yellow);
  background: var(--band-black-soft);
  color: var(--band-yellow);
}

.email-link .pi {
  font-size: 1rem;
  flex-shrink: 0;
}

.social-section {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.social-label {
  font-size: 0.75rem;
  font-weight: 600;
  letter-spacing: 0.1em;
  text-transform: uppercase;
  color: var(--band-white-muted);
}

.embeds {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.embed-wrapper {
  width: 100%;
  overflow: hidden;
  background: var(--band-black-card);
  border: 1px solid var(--color-border);
}

.instagram-embed {
  display: block;
  width: 100%;
  min-height: 480px;
  border: none;
}

.tiktok-wrapper {
  padding: 0.5rem;
}

.reload-btn {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  background: transparent;
  border: 1px solid var(--color-border);
  color: var(--band-white-muted);
  font-size: 0.75rem;
  font-weight: 600;
  letter-spacing: 0.1em;
  text-transform: uppercase;
  padding: 0.5rem 1rem;
  cursor: pointer;
  transition: border-color 0.2s, color 0.2s;
  align-self: flex-start;
}

.reload-btn:hover {
  border-color: var(--band-white-muted);
  color: var(--band-white);
}

/* ── Responsive ── */
@media (max-width: 900px) {
  .contact-layout {
    grid-template-columns: 1fr;
    gap: 3rem;
  }
}

@media (max-width: 768px) {
  .contact-section {
    padding: 3rem 1.5rem 5rem;
  }
}

@media (max-width: 480px) {
  .contact-section {
    padding: 2.5rem 1rem 4rem;
  }
  .contact-layout {
    gap: 2.5rem;
  }
  .email-link {
    font-size: 0.85rem;
  }
}
</style>
