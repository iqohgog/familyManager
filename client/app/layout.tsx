"use client";

import "./globals.css";
import { ThemeProvider } from "@/components/theme-provider";
import { Roboto, Inter } from "next/font/google";
import { ModeTwoggle } from "@/components/theme-menu";
import { Toaster } from "sonner";
const roboto = Roboto({
  variable: "--font-roboto",
  subsets: ["latin", "cyrillic"],
});

const inter = Inter({
  variable: "--font-inter",
  subsets: ["latin", "cyrillic"],
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${roboto.variable}  ${inter.variable}  antialiased relative`}
      >
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <Toaster />
          <ModeTwoggle />
          {children}
        </ThemeProvider>
      </body>
    </html>
  );
}
