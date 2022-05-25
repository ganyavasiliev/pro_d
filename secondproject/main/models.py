from operator import mod
from tabnanny import verbose
from django.db import models


class Text(models.Model):
    title = models.CharField('Название', max_length=50)
    text = models.TextField('Текст')

    def __str__(self):
        return self.title

    class Meta:
        verbose_name = 'Задача'
        verbose_name_plural = 'Задачи'