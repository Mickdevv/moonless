<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { onMounted, reactive, ref } from 'vue';
import { Form } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
import type { RegisterDTO } from '@/types/DTOs/Register.dto';
import router from '@/router';

const authStore = useAuthStore()
const credentials = reactive<RegisterDTO>({
  email: "",
  password1: "",
  password2: ""
})


onMounted(async () => {
  if (await authStore.ensureToken(0, false)) {
    await router.push(`/`)
  }
})

const submitRegister = ($form: any) => {
  authStore.register({
    email: $form.values.email, password1: $form.values.password1, password2:
      $form.values.password2
  })
}


</script>

<template>Register
  <div>
    <Form v-slot="$form" :initialValues="credentials" @submit="submitRegister">
      <div>
        <input-text name="email" type="email" placeholder="Email" fluid />
        <message v-if="$form.email?.invalid" severity=error size="small" variant="simple">{{
          $form.email?.error }}
        </message>
      </div>
      <div>
        <input-text name="password1" type="password" placeholder="Password" fluid />
        <message v-if="$form.password?.invalid" severity=error size="small" variant="simple">{{
          $form.password?.error }}
        </message>
      </div>
      <div>
        <input-text name="password2" type="password" placeholder="Confirm password" fluid />
        <message v-if="$form.password?.invalid" severity=error size="small" variant="simple">{{
          $form.password?.error }}
        </message>
      </div>
      <Button type="submit" severity="secondary" label="Submit" />
    </Form>
  </div>
  <span class="login-suggestion">

    Already have an account? <router-link :to="`/login`">Log in here</router-link>
  </span>
</template>

<style scoped>
.login-suggestion {
  margin-top: 12px;
}

Form {
  gap: 10px;
  display: flex;
  flex-direction: column;
}
</style>
