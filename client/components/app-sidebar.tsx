"use client";

import * as React from "react";
import {
  BookOpen,
  Bot,
  Command,
  Frame,
  LifeBuoy,
  Map,
  PieChart,
  Send,
  Settings2,
  SquareTerminal,
} from "lucide-react";

import { NavProjects } from "@/components/nav-projects";
import { NavUser } from "@/components/nav-user";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";

const fetchUserData = async () => {
  const token = localStorage.getItem("token");
  if (!token) return null;

  try {
    const response = await fetch("http://localhost:8080/iam", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) throw new Error("Failed to fetch user data");

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching user data:", error);
    return null;
  }
};

export function AppSidebar({ ...props }) {
  const [user, setUser] = React.useState({
    name: "",
    email: "",
    avatar: "",
    family: null,
  });
  const [projects, setProjects] = React.useState([]);

  React.useEffect(() => {
    fetchUserData().then((userData) => {
      if (userData) {
        setUser({
          name: `${userData.FirstName} ${userData.LastName}`,
          email: userData.Email,
          avatar: "/avatars/default.jpg",
          family: userData.FamilyID,
        });
        localStorage.setItem("FamilyName", userData.FamilyID);

        if (userData.FamilyID) {
          setProjects([
            { name: "Настройки семьи", url: "#setting", icon: Settings2 },
          ]);
        } else {
          setProjects([
            { name: "Создать семью", url: "#create_family", icon: Frame },
          ]);
        }
      }
    });
  }, []);

  return (
    <Sidebar
      className="top-(--header-height) h-[calc(100svh-var(--header-height))]!"
      {...props}
    >
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" asChild>
              <a href="#">
                <div className="bg-sidebar-primary text-sidebar-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg">
                  <Command className="size-4" />
                </div>
                <div className="grid flex-1 text-left text-sm leading-tight">
                  <span className="truncate font-medium">
                    {user.family ? `Семья: ${user.family}` : "Вы не в семье"}
                  </span>
                  <span className="truncate text-xs"></span>
                </div>
              </a>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <NavProjects projects={projects} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={user} />
      </SidebarFooter>
    </Sidebar>
  );
}
