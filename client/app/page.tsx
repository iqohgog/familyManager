"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/app-sidebar";
import { SiteHeader } from "@/components/site-header";
import { FamilyInvite } from "@/components/family-invite";
import { Settings } from "@/components/settings";
import TaskTable from "@/components/task-table";

export default function Home() {
  const router = useRouter();
  const [activePage, setActivePage] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (!token) {
      router.push("/login");
    }

    if (window.location.hash === "#create_family") {
      setActivePage("createFamily");
    }
    if (window.location.hash === "#setting") {
      setActivePage("setting");
    }

    const handleHashChange = () => {
      if (window.location.hash === "#create_family") {
        setActivePage("createFamily");
      } else if (window.location.hash === "#setting") {
        setActivePage("setting");
      } else {
        setActivePage(null);
      }
    };

    window.addEventListener("hashchange", handleHashChange);

    return () => {
      window.removeEventListener("hashchange", handleHashChange);
    };
  }, [router]);

  return (
    <div className="[--header-height:calc(theme(spacing.14))]">
      <SidebarProvider className="flex flex-col">
        <SiteHeader />
        <div className="flex flex-1">
          <AppSidebar />
          <SidebarInset>
            {activePage === "createFamily" && (
              <div className="flex min-h-svh w-full items-center justify-center p-16 md:p-10">
                <div className="w-full max-w-sm">
                  <FamilyInvite />
                </div>
              </div>
            )}
            {activePage === "setting" && (
              <div className="flex justify-center">
                <div className="w-full max-w-none px-4">
                  <Settings />
                </div>
              </div>
            )}
            {activePage === null &&
              localStorage.getItem("FamilyName") == "null" && (
                <div className="flex min-h-svh w-full items-center justify-center p-16 md:p-10">
                  <div className="w-full max-w-sm">
                    <h1 className="text-xl font-bold">
                      Добро пожаловать в приложение!
                    </h1>
                    <p className="mt-2 text-gray-600">
                      Создайте семью или примите приглашение на почте.
                    </p>
                  </div>
                </div>
              )}
            {activePage === null &&
              localStorage.getItem("FamilyName") != "null" && (
                <div className="flex justify-center">
                  <div className="w-full max-w-none px-4">
                    <TaskTable />
                  </div>
                </div>
              )}
          </SidebarInset>
        </div>
      </SidebarProvider>
    </div>
  );
}
