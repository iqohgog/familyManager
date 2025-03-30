"use client"; // Add this directive at the top

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { RegisterForm } from "@/components/register-form";

export default function Page() {
  const router = useRouter();

  useEffect(() => {
    // Проверяем наличие токена в localStorage (или другом месте, где он может быть)
    const token = localStorage.getItem("token");

    if (token) {
      // Если токен существует, редиректим на главную страницу
      router.push("/");
    }
  }, [router]);

  return (
    <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
      <div className="w-full max-w-sm">
        <RegisterForm />
      </div>
    </div>
  );
}
