"use client"

import { useEffect, useState } from "react"
import { api } from "@/lib/axios"
import { Plus, Search, Edit, Trash2, Heart, Loader2, ImageIcon, Upload, X } from "lucide-react"
import { v4 as uuidv4 } from "uuid"
import { supabase } from "@/lib/supabase"
import { useAuthStore } from "@/store/auth"

import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
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
import { Textarea } from "@/components/ui/textarea"

interface Grammar {
  id: string
  pattern: string
  meaning: string
  example?: string
  note?: string
  source?: string
  jlpt_level?: string
  status?: string
  image_url?: string
  favourite: boolean
}

export default function GrammarPage() {
  const { user } = useAuthStore()
  const [grammars, setGrammars] = useState<Grammar[]>([])
  const [loading, setLoading] = useState(true)
  const [search, setSearch] = useState("")

  const [viewImageUrl, setViewImageUrl] = useState<string | null>(null)

  const [isDialogOpen, setIsDialogOpen] = useState(false)
  const [editId, setEditId] = useState<string | null>(null)
  const [formData, setFormData] = useState({
    pattern: "",
    meaning: "",
    example: "",
    note: "",
    source: "",
    jlpt_level: "",
    image_url: ""
  })
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [isUploading, setIsUploading] = useState(false)
  const [uploadError, setUploadError] = useState("")

  const fetchGrammars = async (query = "") => {
    setLoading(true)
    try {
      const response = await api.get(`/grammars${query ? `?search=${query}` : ""}`)
      setGrammars(response.data.data)
    } catch (err) {
      console.error("Failed to load grammars", err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    const delayDebounceFn = setTimeout(() => {
      fetchGrammars(search)
    }, 500)
    return () => clearTimeout(delayDebounceFn)
  }, [search])

  const toggleFavourite = async (id: string) => {
    try {
      await api.patch(`/grammars/${id}/favourite`)
      setGrammars(grammars.map(g => 
        g.id === id ? { ...g, favourite: !g.favourite } : g
      ))
    } catch (err) {
      console.error("Failed to toggle favourite", err)
    }
  }

  const deleteGrammar = async (id: string) => {
    if (!confirm("Are you sure you want to delete this item?")) return
    
    try {
      await api.delete(`/grammars/${id}`)
      setGrammars(grammars.filter(g => g.id !== id))
    } catch (err) {
      console.error("Failed to delete", err)
    }
  }

  const openEdit = (grammar: Grammar) => {
    setFormData({
      pattern: grammar.pattern,
      meaning: grammar.meaning,
      example: grammar.example || "",
      note: grammar.note || "",
      source: grammar.source || "",
      jlpt_level: grammar.jlpt_level || "",
      image_url: grammar.image_url || ""
    })
    setEditId(grammar.id)
    setIsDialogOpen(true)
  }

  const handleOpenChange = (open: boolean) => {
    setIsDialogOpen(open)
    if (!open) {
      setEditId(null)
      setFormData({ pattern: "", meaning: "", example: "", note: "", source: "", jlpt_level: "", image_url: "" })
      setUploadError("")
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsSubmitting(true)
    try {
      if (editId) {
        await api.put(`/grammars/${editId}`, formData)
      } else {
        await api.post("/grammars", formData)
      }
      setIsDialogOpen(false)
      setEditId(null)
      setFormData({ pattern: "", meaning: "", example: "", note: "", source: "", jlpt_level: "", image_url: "" })
      fetchGrammars(search)
    } catch (err) {
      console.error("Failed to save", err)
    } finally {
      setIsSubmitting(false)
    }
  }

  const handleImageUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (!file) return

    setUploadError("")

    if (file.size > 10 * 1024 * 1024) {
      setUploadError("Image size must be less than 10MB")
      return
    }

    if (!file.type.startsWith('image/')) {
      setUploadError("File must be an image")
      return
    }

    try {
      setIsUploading(true)

      const fileExt = file.name.split('.').pop()
      const fileName = `${user?.id || 'user'}-${uuidv4()}.${fileExt}`
      const filePath = `${fileName}`

      const { error: uploadError } = await supabase.storage
        .from('grammar_notes')
        .upload(filePath, file)

      if (uploadError) {
        throw uploadError
      }

      const { data } = supabase.storage
        .from('grammar_notes')
        .getPublicUrl(filePath)

      setFormData(prev => ({ ...prev, image_url: data.publicUrl }))
    } catch (error: any) {
      setUploadError(error.message || "Failed to upload image")
    } finally {
      setIsUploading(false)
    }
  }

  return (
    <div className="container max-w-6xl mx-auto py-8 px-4 space-y-6">
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Grammar</h1>
          <p className="text-muted-foreground">Manage your Japanese Grammar (Bunpou) collection.</p>
        </div>
        
        <Dialog open={isDialogOpen} onOpenChange={handleOpenChange}>
          <DialogTrigger asChild>
            <Button>
              <Plus className="mr-2 h-4 w-4" /> Add New Grammar
            </Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[500px]">
            <DialogHeader>
              <DialogTitle>{editId ? "Edit Grammar" : "Add Grammar"}</DialogTitle>
              <DialogDescription>
                {editId ? "Edit your Japanese grammar pattern." : "Add a new Japanese grammar pattern to your collection."}
              </DialogDescription>
            </DialogHeader>
            <form onSubmit={handleSubmit} className="space-y-4 py-4">
              <div className="space-y-2">
                <Label htmlFor="pattern">Pattern <span className="text-destructive">*</span></Label>
                <Input 
                  id="pattern" 
                  value={formData.pattern} 
                  onChange={(e) => setFormData({...formData, pattern: e.target.value})} 
                  placeholder="e.g. ～てもいいです"
                  required 
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="meaning">Meaning <span className="text-destructive">*</span></Label>
                <Input 
                  id="meaning" 
                  value={formData.meaning} 
                  onChange={(e) => setFormData({...formData, meaning: e.target.value})} 
                  placeholder="e.g. You may do..."
                  required
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="example">Example Sentence</Label>
                <Textarea 
                  id="example" 
                  value={formData.example} 
                  onChange={(e) => setFormData({...formData, example: e.target.value})} 
                  placeholder="写真を撮ってもいいです。"
                  className="resize-none"
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="note">Usage Notes</Label>
                <Input 
                  id="note" 
                  value={formData.note} 
                  onChange={(e) => setFormData({...formData, note: e.target.value})} 
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="source">Source / JLPT Level</Label>
                <div className="flex gap-2">
                  <Input 
                    id="source" 
                    value={formData.source} 
                    onChange={(e) => setFormData({...formData, source: e.target.value})} 
                    placeholder="Source (e.g. Genki)"
                  />
                  <Input 
                    id="jlpt_level" 
                    value={formData.jlpt_level} 
                    onChange={(e) => setFormData({...formData, jlpt_level: e.target.value})} 
                    placeholder="N5, N4..."
                  />
                </div>
              </div>
              <div className="space-y-2">
                <Label htmlFor="image_upload">Attached Note / Image (Max 10MB)</Label>
                {formData.image_url ? (
                  <div className="relative mt-2 w-full h-40 bg-muted rounded-md overflow-hidden border">
                    <img src={formData.image_url} alt="Note Attachment" className="w-full h-full object-contain" />
                    <Button 
                      type="button" 
                      variant="destructive" 
                      size="icon" 
                      className="absolute top-2 right-2 h-6 w-6 rounded-full"
                      onClick={() => setFormData({...formData, image_url: ""})}
                    >
                      <X className="h-4 w-4" />
                    </Button>
                  </div>
                ) : (
                  <div className="flex items-center gap-4 mt-2">
                    <Input 
                      id="image_upload" 
                      type="file" 
                      accept="image/*" 
                      onChange={handleImageUpload} 
                      disabled={isUploading || isSubmitting}
                      className="cursor-pointer"
                    />
                    {isUploading && <Loader2 className="h-5 w-5 animate-spin text-muted-foreground" />}
                  </div>
                )}
                {uploadError && <p className="text-sm text-destructive">{uploadError}</p>}
              </div>

              <DialogFooter className="mt-6">
                <Button type="submit" disabled={isSubmitting}>
                  {isSubmitting && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
                  Save Item
                </Button>
              </DialogFooter>
            </form>
          </DialogContent>
        </Dialog>
      </div>

      <div className="flex items-center relative max-w-sm">
        <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
        <Input 
          placeholder="Search grammar..." 
          className="pl-9"
          value={search}
          onChange={(e) => setSearch(e.target.value)}
        />
      </div>

      <div className="border rounded-md">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-[50px]"></TableHead>
              <TableHead>Pattern</TableHead>
              <TableHead>Meaning</TableHead>
              <TableHead>Example</TableHead>
              <TableHead className="w-[60px] text-center"><ImageIcon className="h-4 w-4 mx-auto text-muted-foreground"/></TableHead>
              <TableHead>Status</TableHead>
              <TableHead className="text-right">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {loading ? (
              <TableRow>
                <TableCell colSpan={6} className="text-center py-10 text-muted-foreground">
                  <Loader2 className="h-6 w-6 animate-spin mx-auto" />
                </TableCell>
              </TableRow>
            ) : grammars.length === 0 ? (
              <TableRow>
                <TableCell colSpan={7} className="text-center py-10 text-muted-foreground">
                  No Grammar patterns found.
                </TableCell>
              </TableRow>
            ) : (
              grammars.map((item) => (
                <TableRow key={item.id}>
                  <TableCell>
                    <button 
                      onClick={() => toggleFavourite(item.id)}
                      className={`hover:bg-muted p-1.5 rounded-full transition-colors ${item.favourite ? 'text-red-500' : 'text-muted-foreground'}`}
                    >
                      <Heart className="h-4 w-4" fill={item.favourite ? "currentColor" : "none"} />
                    </button>
                  </TableCell>
                  <TableCell className="font-bold text-lg">{item.pattern}</TableCell>
                  <TableCell>{item.meaning}</TableCell>
                  <TableCell className="text-muted-foreground italic truncate max-w-xs">{item.example || "-"}</TableCell>
                  <TableCell className="text-center">
                    {item.image_url && (
                      <button 
                        onClick={() => setViewImageUrl(item.image_url!)} 
                        className="inline-flex hover:bg-muted p-1.5 rounded text-primary transition-colors cursor-pointer"
                        title="View Attached Note"
                      >
                        <ImageIcon className="h-4 w-4" />
                      </button>
                    )}
                  </TableCell>
                  <TableCell>
                    <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-primary/10 text-primary">
                      {item.status || "NEW"}
                    </span>
                  </TableCell>
                  <TableCell className="text-right">
                    <Button variant="ghost" size="icon" title="Edit" onClick={() => openEdit(item)}>
                      <Edit className="h-4 w-4 text-muted-foreground" />
                    </Button>
                    <Button variant="ghost" size="icon" title="Delete" onClick={() => deleteGrammar(item.id)}>
                      <Trash2 className="h-4 w-4 text-destructive" />
                    </Button>
                  </TableCell>
                </TableRow>
              ))
            )}
          </TableBody>
        </Table>
      </div>

      <Dialog open={!!viewImageUrl} onOpenChange={(open) => !open && setViewImageUrl(null)}>
        <DialogContent className="sm:max-w-3xl border-0 p-1 bg-transparent shadow-none">
          <div className="relative w-full rounded-md overflow-hidden bg-black/90 flex items-center justify-center min-h-[300px]">
            {viewImageUrl && (
              <img 
                src={viewImageUrl} 
                alt="Grammar Note Attachment" 
                className="max-w-full max-h-[80vh] object-contain"
              />
            )}
            <Button 
              type="button" 
              variant="secondary" 
              size="icon" 
              className="absolute top-2 right-2 h-8 w-8 rounded-full opacity-70 hover:opacity-100"
              onClick={() => setViewImageUrl(null)}
            >
              <X className="h-4 w-4" />
            </Button>
          </div>
        </DialogContent>
      </Dialog>
    </div>
  )
}
