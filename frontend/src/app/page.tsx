import Link from "next/link";
import {
  ArrowRight,
  BookOpen,
  BrainCircuit,
  Target,
  Flame,
} from "lucide-react";
import { Button } from "@/components/ui/button";

export default function Home() {
  return (
    <div className="flex flex-col min-h-[calc(100vh-3.5rem)]">
      {/* Hero Section */}
      <section
        className="relative flex-1 flex flex-col items-center justify-center text-center px-4 py-20 sm:py-32 bg-cover bg-center bg-no-repeat"
        style={{ backgroundImage: "url('/castle.jpg')" }}
      >
        <div className="absolute inset-0 bg-background/30 dark:bg-background/50 backdrop-blur-[1px]"></div>

        <div className="relative z-10 flex flex-col items-center justify-center w-full max-w-5xl mx-auto">
          <h1 className="text-4xl sm:text-6xl lg:text-7xl font-bold tracking-tight max-w-4xl mb-6 drop-shadow-md">
            Master Japanese with <br className="hidden sm:block" />
            <span className="text-primary drop-shadow-sm">
              Spaced Repetition
            </span>
          </h1>
          <p className="text-xl text-foreground/90 max-w-2xl mb-10 font-medium drop-shadow">
            The ultimate all-in-one tracker for your Japanese learning journey.
            Manage vocabulary, master Kanji, understand grammar, and retain
            knowledge forever.
          </p>
          <div className="flex flex-col sm:flex-row gap-4">
            <Link href="/register">
              <Button
                size="lg"
                className="h-12 px-8 text-lg w-full sm:w-auto shadow-md hover:shadow-lg transition-all"
              >
                Get Started
                <ArrowRight className="ml-2 h-5 w-5" />
              </Button>
            </Link>
            <Link href="/login">
              <Button
                size="lg"
                variant="secondary"
                className="h-12 px-8 text-lg w-full sm:w-auto shadow-sm hover:shadow-md transition-all"
              >
                Sign In
              </Button>
            </Link>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-20 bg-background">
        <div className="container px-4 mx-auto max-w-6xl">
          <div className="text-center mb-16">
            <h2 className="text-3xl font-bold tracking-tight sm:text-4xl mb-4">
              Everything you need to reach fluency
            </h2>
            <p className="text-muted-foreground text-lg max-w-2xl mx-auto">
              Built with scientifically proven learning methods to help you
              remember Japanese forever.
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-8">
            <div className="flex flex-col items-center text-center p-6 bg-card rounded-xl border shadow-sm transition-all hover:shadow-md hover:border-primary/50">
              <div className="h-14 w-14 bg-primary/10 rounded-full flex items-center justify-center mb-6">
                <BrainCircuit className="h-7 w-7 text-primary" />
              </div>
              <h3 className="text-xl font-bold mb-3">Smart SRS Review</h3>
              <p className="text-muted-foreground">
                Our Spaced Repetition System automatically schedules your
                reviews precisely when you're about to forget them.
              </p>
            </div>

            <div className="flex flex-col items-center text-center p-6 bg-card rounded-xl border shadow-sm transition-all hover:shadow-md hover:border-primary/50">
              <div className="h-14 w-14 bg-blue-500/10 rounded-full flex items-center justify-center mb-6">
                <Target className="h-7 w-7 text-blue-500" />
              </div>
              <h3 className="text-xl font-bold mb-3">Goal Tracking</h3>
              <p className="text-muted-foreground">
                Set JLPT targets, track your daily learning streaks, and monitor
                your progress across all language components.
              </p>
            </div>

            <div className="flex flex-col items-center text-center p-6 bg-card rounded-xl border shadow-sm transition-all hover:shadow-md hover:border-primary/50">
              <div className="h-14 w-14 bg-orange-500/10 rounded-full flex items-center justify-center mb-6">
                <Flame className="h-7 w-7 text-orange-500" />
              </div>
              <h3 className="text-xl font-bold mb-3">Comprehensive Logs</h3>
              <p className="text-muted-foreground">
                Build your personal dictionary. Keep detailed notes, readings,
                and meanings for Kanji, Vocab, and Grammar rules.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Footer minimal */}
      <footer className="py-8 border-t text-center text-sm text-muted-foreground">
        <p>
          © {new Date().getFullYear()} Nihongo Learning Journal. All rights
          reserved.
        </p>
      </footer>
    </div>
  );
}
