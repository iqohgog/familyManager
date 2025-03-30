"use client";
import { ModeTwoggle } from "@/components/theme-menu";
import "./globals.css";
import { ThemeProvider } from "@/components/theme-provider";
import { Roboto, Inter } from "next/font/google";

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
      <body className={`${roboto.variable}  ${inter.variable}  antialiased`}>
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <ModeTwoggle />
          {children}
        </ThemeProvider>
      </body>
    </html>
  );
}
