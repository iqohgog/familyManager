"use client";

import * as React from "react";
import { useEffect, useState } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { toast } from "sonner";

export function Settings() {
  const [familyUsers, setFamilyUsers] = useState([]);
  const [inviteEmail, setInviteEmail] = useState("");

  useEffect(() => {
    const fetchFamilyUsers = async () => {
      try {
        const token = localStorage.getItem("token");
        if (!token) return;

        const response = await fetch("http://localhost:8080/family/users", {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        const data = await response.json();
        setFamilyUsers(data);
      } catch (error) {
        console.error("Error fetching family users:", error);
      }
    };
    fetchFamilyUsers();
  }, []);

  const handleInvite = async (e) => {
    e.preventDefault();
    try {
      const token = localStorage.getItem("token");
      if (!token) {
        toast.error("Токен не найден");
        return;
      }

      const response = await fetch("http://localhost:8080/family/invite", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ email: inviteEmail }),
      });

      if (response.ok) {
        toast.success(`Приглашение отправлено на ${inviteEmail}`);
        setInviteEmail("");
      } else {
        toast.error("Ошибка при отправке приглашения");
      }
    } catch (error) {
      console.error("Invite error:", error);
      toast.error("Произошла ошибка");
    }
  };

  const handleExit = async () => {
    if (window.confirm("Вы уверены, что хотите выйти из семьи?")) {
      try {
        const token = localStorage.getItem("token");
        if (!token) {
          toast.error("Токен не найден");
          return;
        }

        const response = await fetch("http://localhost:8080/family/exit", {
          method: "POST",
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (response.ok) {
          toast.success("Вы успешно вышли из семьи");
          // Переадресация на главную страницу с обновлением
          window.location.href = "/";
        } else {
          toast.error("Ошибка при выходе из семьи");
        }
      } catch (error) {
        console.error("Exit error:", error);
        toast.error("Произошла ошибка");
      }
    }
  };

  return (
    <div className="p-4 w-full">
      <h2 className="text-2xl font-bold mb-4">Настройки семьи</h2>
      <div className="w-full overflow-x-auto">
        <h3 className="text-xl font-semibold mb-2">Пользователи семьи</h3>
        <Table className="w-full">
          <TableHeader>
            <TableRow>
              <TableHead className="w-1/4">ID</TableHead>
              <TableHead className="w-1/4">Имя</TableHead>
              <TableHead className="w-1/4">Email</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {familyUsers.map((user) => (
              <TableRow key={user.ID}>
                <TableCell>{user.ID}</TableCell>
                <TableCell>{`${user.FirstName} ${user.LastName}`}</TableCell>
                <TableCell>{user.Email}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
      <div className="mt-6">
        <h3 className="text-xl font-semibold mb-2">Пригласить в семью</h3>
        <form onSubmit={handleInvite} className="flex gap-2">
          <Input
            type="email"
            placeholder="Введите email"
            value={inviteEmail}
            onChange={(e) => setInviteEmail(e.target.value)}
            className="max-w-md"
          />
          <Button type="submit">Пригласить</Button>
        </form>
      </div>
      <div className="mt-6">
        <h3 className="text-xl font-semibold mb-2">Выйти из семьи</h3>
        <Button onClick={handleExit}>Выйти</Button>
      </div>
    </div>
  );
}
