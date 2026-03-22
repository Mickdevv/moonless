<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import { reactive, ref } from 'vue';
import { Form } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
import type { RegisterDTO } from '@/types/DTOs/Register.dto';

const authStore = useAuthStore()
const credentials = reactive<RegisterDTO>({
  email: "",
  password1: "",
  password2: ""
})


const submitRegister = ($form: any) => {
  authStore.register({
    email: $form.values.email, password1: $form.values.password1, password2:
      $form.values.password2
  })
}


</script>

<template>Register screen
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
  <router-link :to="`/register`" class="editButton
            button">Register</router-link>
</template>

<style scoped></style>
