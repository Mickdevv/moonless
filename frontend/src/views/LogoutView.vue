<script setup lang="ts">
import { useAuthStore } from '@/stores/auth';
import type { LoginDTO } from '@/types/DTOs/Login.dto';
import { onMounted, reactive, ref } from 'vue';
import { Form } from '@primevue/forms';
import { InputText, Message, Button } from 'primevue';
import router from '@/router';

const authStore = useAuthStore()
const credentials = reactive<LoginDTO>({
  email: "",
  password: ""
})

await authStore.logout()
onMounted(async () => {
})

const submitLogin = ($form: any) => {
  authStore.login({ email: $form.values.email, password: $form.values.password })
}


</script>

<template>Log in
  <!-- <div> -->
  <!--   <Form v-slot="$form" :initialValues="credentials" @submit="submitLogin"> -->
  <!--     <div> -->
  <!--       <input-text name="email" type="email" placeholder="Email" fluid /> -->
  <!--       <message v-if="$form.email?.invalid" severity=error size="small" variant="simple">{{ -->
  <!--         $form.email?.error }} -->
  <!--       </message> -->
  <!--     </div> -->
  <!--     <div> -->
  <!--       <input-text name="password" type="password" placeholder="Password" fluid /> -->
  <!--       <message v-if="$form.password?.invalid" severity=error size="small" variant="simple">{{ -->
  <!--         $form.password?.error }} -->
  <!--       </message> -->
  <!--     </div> -->
  <!--     <Button type="submit" severity="secondary" label="Submit" /> -->
  <!--   </Form> -->
  <!-- </div> -->
  <span class="register-suggestion">
    You are low logged out
  </span>
</template>

<style scoped>
.register-suggestion {
  margin-top: 12px;
}

Form {
  gap: 10px;
  display: flex;
  flex-direction: column;
}
</style>
