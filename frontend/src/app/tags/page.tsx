"use client"

import { useEffect, useState } from "react"
import { api } from "@/lib/axios"
import { Plus, Edit, Trash2, Loader2, Tag, Palette } from "lucide-react"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog"
import { Label } from "@/components/ui/label"

interface TagItem {
  id: string
  name: string
  color?: string | null
  createdAt: string
  updatedAt: string
}

const PRESET_COLORS = [
  "#ef4444", // red
  "#f97316", // orange
  "#f59e0b", // amber
  "#eab308", // yellow
  "#84cc16", // lime
  "#22c55e", // green
  "#14b8a6", // teal
  "#06b6d4", // cyan
  "#3b82f6", // blue
  "#6366f1", // indigo
  "#8b5cf6", // violet
  "#a855f7", // purple
  "#d946ef", // fuchsia
  "#ec4899", // pink
  "#f43f5e", // rose
  "#78716c", // stone
]

export default function TagsPage() {
  const [tags, setTags] = useState<TagItem[]>([])
  const [loading, setLoading] = useState(true)

  const [isDialogOpen, setIsDialogOpen] = useState(false)
  const [editId, setEditId] = useState<string | null>(null)
  const [formData, setFormData] = useState({
    name: "",
    color: "#3b82f6",
  })
  const [isSubmitting, setIsSubmitting] = useState(false)

  const fetchTags = async () => {
    setLoading(true)
    try {
      const response = await api.get("/tags")
      setTags(response.data.data || [])
    } catch (err) {
      console.error("Failed to load tags", err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchTags()
  }, [])

  const deleteTag = async (id: string) => {
    if (!confirm("Are you sure you want to delete this tag? It will be removed from all items.")) return

    try {
      await api.delete(`/tags/${id}`)
      setTags(tags.filter(t => t.id !== id))
    } catch (err) {
      console.error("Failed to delete tag", err)
    }
  }

  const openEdit = (tag: TagItem) => {
    setFormData({
      name: tag.name,
      color: tag.color || "#3b82f6",
    })
    setEditId(tag.id)
    setIsDialogOpen(true)
  }

  const handleOpenChange = (open: boolean) => {
    setIsDialogOpen(open)
    if (!open) {
      setEditId(null)
      setFormData({ name: "", color: "#3b82f6" })
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsSubmitting(true)
    try {
      const payload = {
        name: formData.name,
        color: formData.color || null,
      }

      if (editId) {
        await api.put(`/tags/${editId}`, payload)
      } else {
        await api.post("/tags", payload)
      }
      setIsDialogOpen(false)
      setEditId(null)
      setFormData({ name: "", color: "#3b82f6" })
      fetchTags()
    } catch (err) {
      console.error("Failed to save tag", err)
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <div className="container max-w-6xl mx-auto py-8 px-4 space-y-6">
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Tags</h1>
          <p className="text-muted-foreground">Organize your vocabulary, kanji, and grammar with tags.</p>
        </div>

        <Dialog open={isDialogOpen} onOpenChange={handleOpenChange}>
          <DialogTrigger asChild>
            <Button>
              <Plus className="mr-2 h-4 w-4" /> New Tag
            </Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[420px]">
            <DialogHeader>
              <DialogTitle>{editId ? "Edit Tag" : "Create Tag"}</DialogTitle>
              <DialogDescription>
                {editId ? "Update your tag details." : "Create a new tag to organize your study items."}
              </DialogDescription>
            </DialogHeader>
            <form onSubmit={handleSubmit} className="space-y-5 py-4">
              <div className="space-y-2">
                <Label htmlFor="tag-name">Name <span className="text-destructive">*</span></Label>
                <Input
                  id="tag-name"
                  value={formData.name}
                  onChange={(e) => setFormData({ ...formData, name: e.target.value })}
                  placeholder="e.g. N5, Daily Life, Anime"
                  required
                  maxLength={100}
                />
              </div>
              <div className="space-y-3">
                <Label className="flex items-center gap-2">
                  <Palette className="h-4 w-4" /> Color
                </Label>
                <div className="grid grid-cols-8 gap-2">
                  {PRESET_COLORS.map((color) => (
                    <button
                      key={color}
                      type="button"
                      onClick={() => setFormData({ ...formData, color })}
                      className="h-8 w-8 rounded-full border-2 transition-all hover:scale-110 focus:outline-none"
                      style={{
                        backgroundColor: color,
                        borderColor: formData.color === color ? "var(--foreground)" : "transparent",
                        boxShadow: formData.color === color ? `0 0 0 2px var(--background), 0 0 0 4px ${color}` : "none",
                      }}
                      title={color}
                    />
                  ))}
                </div>
                <div className="flex items-center gap-3 pt-1">
                  <Label htmlFor="custom-color" className="text-sm text-muted-foreground whitespace-nowrap">Custom:</Label>
                  <div className="flex items-center gap-2 flex-1">
                    <input
                      type="color"
                      id="custom-color"
                      value={formData.color}
                      onChange={(e) => setFormData({ ...formData, color: e.target.value })}
                      className="h-8 w-8 rounded cursor-pointer border border-border"
                    />
                    <Input
                      value={formData.color}
                      onChange={(e) => setFormData({ ...formData, color: e.target.value })}
                      className="font-mono text-sm h-8"
                      placeholder="#000000"
                      maxLength={7}
                    />
                  </div>
                </div>
              </div>
              <DialogFooter>
                <Button type="submit" disabled={isSubmitting}>
                  {isSubmitting && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                  {editId ? "Update Tag" : "Create Tag"}
                </Button>
              </DialogFooter>
            </form>
          </DialogContent>
        </Dialog>
      </div>

      {loading ? (
        <div className="flex items-center justify-center py-20">
          <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
        </div>
      ) : tags.length === 0 ? (
        <div className="flex flex-col items-center justify-center py-20 text-center space-y-4">
          <div className="h-16 w-16 rounded-full bg-muted flex items-center justify-center">
            <Tag className="h-8 w-8 text-muted-foreground" />
          </div>
          <div>
            <h3 className="text-lg font-semibold">No tags yet</h3>
            <p className="text-muted-foreground text-sm mt-1">Create your first tag to start organizing your study items.</p>
          </div>
        </div>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
          {tags.map((tag) => (
            <div
              key={tag.id}
              className="group relative border rounded-lg p-4 transition-all hover:shadow-md hover:border-primary/30 bg-card"
            >
              <div className="flex items-start justify-between gap-2">
                <div className="flex items-center gap-3 min-w-0 flex-1">
                  <div
                    className="h-4 w-4 rounded-full shrink-0 ring-2 ring-offset-2 ring-offset-background"
                    style={{
                      backgroundColor: tag.color || "#78716c",
                      ringColor: tag.color || "#78716c",
                    }}
                  />
                  <span className="font-medium truncate">{tag.name}</span>
                </div>
                <div className="flex items-center gap-0.5 opacity-0 group-hover:opacity-100 transition-opacity">
                  <Button
                    variant="ghost"
                    size="icon"
                    className="h-8 w-8"
                    title="Edit"
                    onClick={() => openEdit(tag)}
                  >
                    <Edit className="h-3.5 w-3.5 text-muted-foreground" />
                  </Button>
                  <Button
                    variant="ghost"
                    size="icon"
                    className="h-8 w-8"
                    title="Delete"
                    onClick={() => deleteTag(tag.id)}
                  >
                    <Trash2 className="h-3.5 w-3.5 text-destructive" />
                  </Button>
                </div>
              </div>
              <p className="text-xs text-muted-foreground mt-2">
                Created {new Date(tag.createdAt).toLocaleDateString("en-US", { month: "short", day: "numeric", year: "numeric" })}
              </p>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
