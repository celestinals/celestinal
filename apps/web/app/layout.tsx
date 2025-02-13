import { Geist, Geist_Mono } from "next/font/google";

import "@tickex/ui/globals.css";
import { Providers } from "@/components/providers";
import { Metadata } from "next";
import React from "react";

const fontSans = Geist({
  subsets: ["latin"],
  variable: "--font-sans",
});

const fontMono = Geist_Mono({
  subsets: ["latin"],
  variable: "--font-mono",
});

export const metadata: Metadata = {
  title: "Tickex",
  description:
    "Buying, selling, exchanging, and sharing all types of tickets and game cards on a secure cloud-native platform.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body
        className={`${fontSans.variable} ${fontMono.variable} font-sans antialiased `}
      >
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
