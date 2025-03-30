"use client";

import { useState } from "react";
import { useRouter } from "next/navigation"; // Для использования навигации
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  CardDescription,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

export function FamilyInvite({
  className,
  ...props
}: React.ComponentProps<"div">) {
  const [name, setName] = useState("");
  const router = useRouter(); // Для использования навигации

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();

    const token = localStorage.getItem("token"); // Получаем токен из localStorage

    if (!token) {
      console.error("Токен не найден в localStorage");
      return;
    }

    try {
      const response = await fetch("http://localhost:8080/family/create", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`, // Добавляем токен в заголовки
        },
        body: JSON.stringify({ name }),
      });

      if (response.ok) {
        // Перенаправляем на главную страницу
        router.replace("/");

        // Даем небольшую задержку перед обновлением страницы, чтобы редирект успел отработать
        setTimeout(() => {
          window.location.reload();
        }, 100);
      } else {
        console.error("Ошибка при создании семьи");
      }
    } catch (error) {
      console.error("Произошла ошибка:", error);
    }
  };

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader>
          <CardTitle>Создайте семью</CardTitle>
          <CardDescription>Введите данные для создания семьи</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <div className="flex flex-col gap-6">
              <div className="grid gap-3">
                <Label htmlFor="name">Название семьи</Label>
                <Input
                  id="name"
                  type="text"
                  required
                  value={name}
                  onChange={(e) => setName(e.target.value)} // Обновляем состояние
                />
              </div>
              <Button type="submit" className="w-full">
                Создать семью
              </Button>
            </div>
            <div className="mt-4 text-center text-sm">
              У вас уже есть семья?{" "}
              <a href="#create_family" className="underline underline-offset-4">
                <br />
                Попросите создателя пригласить вас
              </a>
            </div>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
