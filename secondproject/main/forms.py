from matplotlib import widgets
from .models import Text
from django.forms import ModelForm, TextInput, Textarea


class TextForm(ModelForm):
    class Meta:
        model = Text
        fields = ["title", "text"]
        widgets = {
            "title": TextInput(attrs={
                'class': 'form-control',
                'placeholder': 'Введите название'
                }),
            "text": Textarea(attrs={
                'class': 'form-control',
                'placeholder': 'Введите описание'
                }),
            }

