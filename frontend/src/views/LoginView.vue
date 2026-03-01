<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import type { LoginDTO } from '@/types/DTOs/Login.dto';
import { reactive, ref } from 'vue';
import { Form } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';

const authStore = useAuthStore()
const credentials = reactive<LoginDTO>({
  email: "",
  password: ""
})


const submitLogin = ($form: any) => {
  authStore.login({ email: $form.values.email, password: $form.values.password })
}


</script>

<template>Login screen
  <div>
    <Form v-slot="$form" :initialValues="credentials" @submit="submitLogin">
      <div>
        <input-text name="email" type="email" placeholder="Email" fluid />
        <message v-if="$form.email?.invalid" severity=error size="small" variant="simple">{{
          $form.email?.error }}
        </message>
      </div>
      <div>
        <input-text name="password" type="password" placeholder="Password" fluid />
        <message v-if="$form.password?.invalid" severity=error size="small" variant="simple">{{
          $form.password?.error }}
        </message>
      </div>
      <Button type="submit" severity="secondary" label="Submit" />
    </Form>
  </div>
  <router-link :to="`/register`" class="editButton
            button">Register</router-link>
</template>

<style scoped></style>
